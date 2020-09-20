package gl

import (
  "fmt"
  "log"

  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/glfw/v3.3/glfw"
)

func Init() error {
  if err := glfw.Init(); err != nil {
    log.Fatalln("Failed to initialize glfw:", err)
    return err
  }
  glfw.WindowHint(glfw.ContextVersionMajor, 4)
  glfw.WindowHint(glfw.ContextVersionMinor, 1)
  glfw.WindowHint(glfw.Resizable, glfw.False)
  glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
  glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
  if err := gl.Init(); err != nil {
    log.Fatalln("Failed to initialize window:", err)
    return err
  }
  version := gl.GoStr(gl.GetString(gl.VERSION))
  fmt.Println("OpenGL version", version)
  return nil
}
