package gl

import (
  "github.com/go-gl/mathgl/mgl32"
)

type Model struct {
  transform mgl32.Mat4

  vao *VAO

  texture *Texture
}
