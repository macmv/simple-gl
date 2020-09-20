package gl

import (
  "github.com/go-gl/mathgl/mgl32"
)

func NewCube(width, height, depth float32) *Model {
  indices := []int32{
    // Bottom
    0, 1, 2,
    3, 2, 1,

    // Top
    6, 7, 5,
    6, 5, 4,

    // Front
    0, 4, 5,
    0, 5, 1,

    // Back
    2, 3, 6,
    7, 6, 3,

    // Left
    3, 5, 7,
    5, 3, 1,

    // Right
    6, 4, 2,
    0, 2, 4,
  }
  verts := []float32{
    // Bottom
    0, 0, 0,
    1, 0, 0,
    0, 0, 1,
    1, 0, 1,
    // Top
    0, 1, 0,
    1, 1, 0,
    0, 1, 1,
    1, 1, 1,
  }
  uvs := []float32{
    // Bottom
    0, 0,
    1, 0,
    0, 1,
    1, 1,
    // Top
    0, 0,
    1, 0,
    0, 1,
    1, 1,
  }
  normals := []float32{
    // Bottom
    0, 0, 0,
    0, 0, 0,
    0, 0, 0,
    0, 0, 0,
    // Top
    0, 0, 0,
    0, 0, 0,
    0, 0, 0,
    0, 0, 0,
  }
  vao := NewVAO()
  vao.SetData(indices, verts, uvs, normals)
  m := Model{}
  m.vao = vao
  m.Transform = mgl32.Ident4()
  return &m
}
