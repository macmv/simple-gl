package core

import (
  "github.com/go-gl/mathgl/mgl32"
)

type Shader interface {
  LoadPerspective(win Window, near, far float32)
  LoadCamera(x, y, z float32)

  StoreUniformMat4f(name string, mat mgl32.Mat4)
  StoreUniform3f(name string, vec mgl32.Vec3)

  Id() uint32
}
