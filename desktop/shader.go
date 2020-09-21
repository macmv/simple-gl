package desktop

import (
  "fmt"
  "strings"
  "io/ioutil"

  "github.com/macmv/simple-gl/core"

  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/mathgl/mgl32"
)

type Shader struct {
  id uint32 
  locations map[string]int32
}

func (c *Core) NewShaderGeo(geometry_path, vertex_path, fragment_path string) (core.Shader, error) {
  var geometry_source []byte
  if geometry_path != "" {
    var err error
    geometry_source, err = ioutil.ReadFile(geometry_path)
    if err != nil {
      return nil, err
    }
  }
  vertex_source, err := ioutil.ReadFile(vertex_path)
  if err != nil {
    return nil, err
  }
  fragment_source, err := ioutil.ReadFile(fragment_path)
  if err != nil {
    return nil, err
  }

  var geometry uint32
  if geometry_path != "" {
    geometry, err = compile_shader(string(geometry_source) + "\x00", gl.GEOMETRY_SHADER)
    if err != nil {
      return nil, err
    }
  }
  vertex, err := compile_shader(string(vertex_source) + "\x00", gl.VERTEX_SHADER)
  if err != nil {
    return nil, err
  }
  fragment, err := compile_shader(string(fragment_source) + "\x00", gl.FRAGMENT_SHADER)
  if err != nil {
    return nil, err
  }

  program := gl.CreateProgram()

  if geometry_path != "" {
    gl.AttachShader(program, geometry)
  }
  gl.AttachShader(program, vertex)
  gl.AttachShader(program, fragment)
  gl.LinkProgram(program)

  var status int32
  gl.GetProgramiv(program, gl.LINK_STATUS, &status)

  if status == gl.FALSE {
    var logLength int32
    gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

    log := strings.Repeat("\x00", int(logLength+1))
    gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

    return nil, fmt.Errorf("Failed to link program: %v", log)
  }

  s := Shader{}
  s.id = program
  s.locations = make(map[string]int32)
  (&s).LoadUniform("projection")
  (&s).LoadUniform("camera")
  (&s).LoadUniform("model")
  (&s).LoadUniform("color")
  return &s, nil
}

func (c *Core) NewShader(vertex_path, fragment_path string) (core.Shader, error) {
  return c.NewShaderGeo("", vertex_path, fragment_path)
}


func compile_shader(source string, stype uint32) (uint32, error) {
  shader := gl.CreateShader(stype)
  csource, free := gl.Strs(source)
  gl.ShaderSource(shader, 1, csource, nil)
  free()
  gl.CompileShader(shader)

  var status int32
  gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
  if status == gl.FALSE {
    var logLength int32
    gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

    log := strings.Repeat("\x00", int(logLength+1))
    gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

    return 0, fmt.Errorf("failed to compile %v: %v", source, log)
  }

  return shader, nil
}

func (s *Shader) LoadUniform(name string) {
  _, ok := s.locations[name]
  if !ok {
    loc := gl.GetUniformLocation(s.id, gl.Str(name + "\x00"))
    code := gl.GetError()
    if code != 0 {
      fmt.Println("Got loc:", loc)
      fmt.Println("Name:", name)
      fmt.Println("Code:", code)
      panic("Got error while getting uniform location")
    }
    s.locations[name] = loc
  }
}

func (s *Shader) StoreUniformMat4f(name string, val mgl32.Mat4) {
  loc, ok := s.locations[name]
  if !ok {
    panic("Invalid name " + name)
  }
  gl.UniformMatrix4fv(loc, 1, false, &val[0])
}

func (s *Shader) StoreUniform3f(name string, val mgl32.Vec3) {
  loc, ok := s.locations[name]
  if !ok {
    panic("Invalid name " + name)
  }
  gl.Uniform3f(loc, val[0], val[1], val[2])
}

func (s *Shader) LoadPerspective(window core.Window, near, far float32) {
  projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(window.Width()) / float32(window.Height()), near, far)
  s.StoreUniformMat4f("projection", projection)
}

func (s *Shader) LoadCamera(x, y, z float32) {
  camera := mgl32.LookAtV(mgl32.Vec3{x, y, z}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
  s.StoreUniformMat4f("camera", camera)
}

func (s *Shader) Id() uint32 {
  return s.id
}
