package gl

import (
  "github.com/go-gl/mathgl/mgl32"
)

type Model struct {
  Transform mgl32.Mat4
  Color mgl32.Vec3

  vao *VAO

  texture *Texture
}

func (m *Model) Vao() *VAO {
  return m.vao
}
