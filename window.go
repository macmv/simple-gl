package gl

import (
  "runtime"

  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/glfw/v3.3/glfw"
)

type Window struct {
  width, height int
  glfw *glfw.Window
  current_shader *Shader
}

func NewWindow(title string, width, height int) *Window {
  glfw_window, err := glfw.CreateWindow(width, height, title, nil, nil)
  if err != nil {
    panic(err)
  }
  glfw_window.MakeContextCurrent()
  gl.Enable(gl.CULL_FACE)
  gl.Enable(gl.DEPTH_TEST)

  w := Window{}
  w.width = width
  w.height = height
  w.glfw = glfw_window
  return &w
}

func (w *Window) Use(shader *Shader) {
  if w.current_shader == nil {
    gl.UseProgram(shader.id)
    w.current_shader = shader
  }
}

func (w *Window) Render(model *Model) error {
  if w.current_shader == nil {
    return nil
  }
  w.current_shader.StoreUniformMat4f("model", model.Transform)
  w.current_shader.StoreUniform3f("color", model.Color)

  // gl.ActiveTexture(gl.TEXTURE0)
  // model.texture.Bind()
  model.vao.Bind()
  gl.DrawElements(gl.TRIANGLES, model.vao.length, gl.UNSIGNED_INT, nil)
  model.vao.Unbind()

  return nil
}

func (w *Window) Finish() {
  w.current_shader = nil
}

func (w *Window) Sync() {
  w.glfw.SwapBuffers()
  glfw.PollEvents()
  gl.ClearColor(0, 0, 0, 1);
  gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT);
}

func (w *Window) Closed() bool {
  return w.glfw.ShouldClose()
}

func (w *Window) Close() {
  glfw.Terminate()
}

func (w *Window) Width() int { return w.width }
func (w *Window) Height() int { return w.height }

func init() {
  // GLFW event handling must run on the main OS thread
  runtime.LockOSThread()
}

func main() {
  // projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
  // gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])
  // 
  // camera := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
  // cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
  // gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])
  // 
  // model := mgl32.Ident4()
  // modelUniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
  // gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])
  // 
  // textureUniform := gl.GetUniformLocation(program, gl.Str("tex\x00"))
  // gl.Uniform1i(textureUniform, 0)
  // 
  // gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))
  // 
  // // Load the texture
  // texture, err := newTexture("square.png")
  // if err != nil {
  //   log.Fatalln(err)
  // }
  // 
  // // Configure the vertex data
  // var vao uint32
  // gl.GenVertexArrays(1, &vao)
  // gl.BindVertexArray(vao)
  // 
  // var vbo uint32
  // gl.GenBuffers(1, &vbo)
  // gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
  // gl.BufferData(gl.ARRAY_BUFFER, len(cubeVertices)*4, gl.Ptr(cubeVertices), gl.STATIC_DRAW)
  // 
  // vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
  // gl.EnableVertexAttribArray(vertAttrib)
  // gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))
  // 
  // texCoordAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertTexCoord\x00")))
  // gl.EnableVertexAttribArray(texCoordAttrib)
  // gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
  // 
  // // Configure global settings
  // gl.Enable(gl.DEPTH_TEST)
  // gl.DepthFunc(gl.LESS)
  // gl.ClearColor(1.0, 1.0, 1.0, 1.0)
  // 
  // angle := 0.0
  // previousTime := glfw.GetTime()
  // 
  // for !window.ShouldClose() {
  //   gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
  // 
  //   // Update
  //   time := glfw.GetTime()
  //   elapsed := time - previousTime
  //   previousTime = time
  // 
  //   angle += elapsed
  //   model = mgl32.HomogRotate3D(float32(angle), mgl32.Vec3{0, 1, 0})
  // 
  //   // Render
  // }
}

// func newTexture(file string) (uint32, error) {
//   imgFile, err := os.Open(file)
//   if err != nil {
//     return 0, fmt.Errorf("texture %q not found on disk: %v", file, err)
//   }
//   img, _, err := image.Decode(imgFile)
//   if err != nil {
//     return 0, err
//   }
// 
//   rgba := image.NewRGBA(img.Bounds())
//   if rgba.Stride != rgba.Rect.Size().X*4 {
//     return 0, fmt.Errorf("unsupported stride")
//   }
//   draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
// 
//   var texture uint32
//   gl.GenTextures(1, &texture)
//   gl.ActiveTexture(gl.TEXTURE0)
//   gl.BindTexture(gl.TEXTURE_2D, texture)
//   gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
//   gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
//   gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
//   gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
//   gl.TexImage2D(
//     gl.TEXTURE_2D,
//     0,
//     gl.RGBA,
//     int32(rgba.Rect.Size().X),
//     int32(rgba.Rect.Size().Y),
//     0,
//     gl.RGBA,
//     gl.UNSIGNED_BYTE,
//     gl.Ptr(rgba.Pix)
//   )
// 
//   return texture, nil
// }
// 
// // Set the working directory to the root of Go package, so that its assets can be accessed.
// func init() {
//   dir, err := importPathToDir("github.com/go-gl/example/gl41core-cube")
//   if err != nil {
//     log.Fatalln("Unable to find Go package in your GOPATH, it's needed to load assets:", err)
//   }
//   err = os.Chdir(dir)
//   if err != nil {
//     log.Panicln("os.Chdir:", err)
//   }
// }
// 
// // importPathToDir resolves the absolute path from importPath.
// // There doesn't need to be a valid Go package inside that import path,
// // but the directory must exist.
// func importPathToDir(importPath string) (string, error) {
//   p, err := build.Import(importPath, "", build.FindOnly)
//   if err != nil {
//     return "", err
//   }
//   return p.Dir, nil
// }
