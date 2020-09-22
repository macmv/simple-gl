package core

import (
  "fmt"
  "io/ioutil"

  "golang.org/x/mobile/gl"

  "github.com/go-gl/mathgl/mgl32"
)

type Shader struct {
  gl gl.Context3
  program gl.Program
  locations map[string]gl.Uniform
}

func NewShader(glctx gl.Context3, vertex_path, fragment_path string) (*Shader, error) {
  vertex_source, err := ioutil.ReadFile(vertex_path)
  if err != nil {
    return nil, err
  }
  fragment_source, err := ioutil.ReadFile(fragment_path)
  if err != nil {
    return nil, err
  }

  vertex, err := compile_shader(glctx, string(vertex_source), gl.VERTEX_SHADER)
  if err != nil {
    return nil, err
  }
  fragment, err := compile_shader(glctx, string(fragment_source), gl.FRAGMENT_SHADER)
  if err != nil {
    return nil, err
  }

  program := glctx.CreateProgram()

  glctx.AttachShader(program, vertex)
  glctx.AttachShader(program, fragment)
  glctx.LinkProgram(program)

  status := glctx.GetProgrami(program, gl.LINK_STATUS)

  if status == gl.FALSE {
    log := glctx.GetProgramInfoLog(program)

    return nil, fmt.Errorf("Failed to link program: %v", log)
  }

  s := Shader{}
  s.gl = glctx
  s.program = program
  s.locations = make(map[string]gl.Uniform)
  (&s).LoadUniform("projection")
  (&s).LoadUniform("camera")
  (&s).LoadUniform("model")
  (&s).LoadUniform("color")
  return &s, nil
}


func compile_shader(glctx gl.Context, source string, stype gl.Enum) (gl.Shader, error) {
  shader := glctx.CreateShader(stype)
  glctx.ShaderSource(shader, source)
  glctx.CompileShader(shader)

  status := glctx.GetShaderi(shader, gl.COMPILE_STATUS)
  if status == gl.FALSE {
    log := glctx.GetShaderInfoLog(shader)

    return shader, fmt.Errorf("failed to compile %v: %v", source, log)
  }

  return shader, nil
}

func (s *Shader) LoadUniform(name string) {
  _, ok := s.locations[name]
  if !ok {
    loc := s.gl.GetUniformLocation(s.program, name)
    code := s.gl.GetError()
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
  s.gl.UniformMatrix4fv(loc, val[:])
}

func (s *Shader) StoreUniform3f(name string, val mgl32.Vec3) {
  loc, ok := s.locations[name]
  if !ok {
    panic("Invalid name " + name)
  }
  s.gl.Uniform3f(loc, val[0], val[1], val[2])
}

func (s *Shader) LoadPerspective(window Window, near, far float32) {
  projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(window.Width()) / float32(window.Height()), near, far)
  s.StoreUniformMat4f("projection", projection)
}

func (s *Shader) LoadCamera(x, y, z float32) {
  camera := mgl32.LookAtV(mgl32.Vec3{x, y, z}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
  s.StoreUniformMat4f("camera", camera)
}

func (s *Shader) Program() gl.Program {
  return s.program
}
