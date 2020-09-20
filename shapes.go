package gl

import (
  "github.com/go-gl/mathgl/mgl32"
)

func NewCube(width, height, depth float32) *Model {
  indices := []int32{
    // Bottom
    2, 1, 0,
    2, 1, 3,
  }
  verts := []float32{
    // Bottom
    0, 0, 0,
    1, 0, 0,
    0, 1, 1,
    1, 1, 1,
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
  m.transform = mgl32.Ident4()
  return &m
}
