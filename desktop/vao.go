package desktop

import (
  "github.com/macmv/simple-gl/core"

  "github.com/go-gl/gl/v4.1-core/gl"
)

type VAO struct {
  id uint32
  length int32
}

func (c *Core) NewVAO() core.VAO {
  var id uint32
  gl.GenVertexArrays(1, &id)

  v := VAO{}
  v.id = id
  return &v
}

func (v *VAO) Length() int32 {
  return v.length
}

func (v *VAO) Bind() {
  gl.BindVertexArray(v.id)
  gl.EnableVertexAttribArray(0)
  gl.EnableVertexAttribArray(1)
  gl.EnableVertexAttribArray(2)
}

func (v *VAO) Unbind() {
  gl.DisableVertexAttribArray(0)
  gl.DisableVertexAttribArray(1)
  gl.DisableVertexAttribArray(2)
  gl.BindVertexArray(0)
}

func (v *VAO) SetData(indices []int32, vertices, uvs, normals []float32) {
  v.length = int32(len(indices))

  // bind the vao
  gl.BindVertexArray(v.id)

  var vbo uint32
  // create empty buffer
  gl.GenBuffers(1, &vbo)
  // bind the new buffer
  gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, vbo)
  // 4 is sizeof float
  gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices) * 4, gl.Ptr(indices), gl.STATIC_DRAW)
  // automatically bound to vao

  // create empty buffer
  gl.GenBuffers(1, &vbo)
  // bind the new buffer
  gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
  // set the data (4 is sizeof float)
  gl.BufferData(gl.ARRAY_BUFFER, len(vertices) * 4, gl.Ptr(vertices), gl.STATIC_DRAW)
  // bind vbo to vao
  gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
  // unbind vbo
  gl.BindBuffer(gl.ARRAY_BUFFER, 0)

  // create empty buffer
  gl.GenBuffers(1, &vbo)
  // bind the new buffer
  gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
  // set the data (4 is sizeof float)
  gl.BufferData(gl.ARRAY_BUFFER, len(uvs) * 4, gl.Ptr(uvs), gl.STATIC_DRAW)
  // bind vbo to vao
  gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 0, nil)
  // unbind vbo
  gl.BindBuffer(gl.ARRAY_BUFFER, 0)

  // create empty buffer
  gl.GenBuffers(1, &vbo)
  // bind the new buffer
  gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
  // set the data (4 is sizeof float)
  gl.BufferData(gl.ARRAY_BUFFER, len(normals) * 4, gl.Ptr(normals), gl.STATIC_DRAW)
  // bind vbo to vao
  gl.VertexAttribPointer(2, 3, gl.FLOAT, false, 0, nil)
  // unbind vbo
  gl.BindBuffer(gl.ARRAY_BUFFER, 0)

  // unbind vao
  gl.BindVertexArray(0)
}


