package desktop

import (
  "math"
  "golang.org/x/mobile/gl"

  "github.com/macmv/simple-gl/core"
)

type VAO struct {
  va gl.VertexArray
  gl gl.Context
  length int
}

func (c *Core) NewVAO() core.VAO {
  v := VAO{}
  v.va = c.Gl().CreateVertexArray()
  v.gl = c.Gl()
  return &v
}

func (v *VAO) Length() int {
  return v.length
}

func (v *VAO) Bind() {
  v.gl.BindVertexArray(v.va)
  // v.gl.EnableVertexAttribArray(gl.Attrib{0})
  // v.gl.EnableVertexAttribArray(gl.Attrib{1})
  // v.gl.EnableVertexAttribArray(gl.Attrib{2})
}

func (v *VAO) Unbind() {
  // v.gl.DisableVertexAttribArray(gl.Attrib{0})
  // v.gl.DisableVertexAttribArray(gl.Attrib{1})
  // v.gl.DisableVertexAttribArray(gl.Attrib{2})
  v.gl.BindVertexArray(v.va)
}

func (v *VAO) SetData(indices []int32, vertices, uvs, normals []float32) {
  v.length = len(indices)

  // bind the vao
  v.gl.BindVertexArray(v.va)

  // create empty buffer
  vbo := v.gl.CreateBuffer()
  // bind the new buffer
  v.gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, vbo)
  // set data in buffer
  v.gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, int_array_to_byte_array(indices), gl.STATIC_DRAW)
  // automatically bound to vao

  // create empty buffer
  vbo = v.gl.CreateBuffer()
  // bind the new buffer
  v.gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
  // set the data (4 is sizeof float)
  v.gl.BufferData(gl.ARRAY_BUFFER, float_array_to_byte_array(vertices), gl.STATIC_DRAW)
  // bind vbo to shader
  // v.gl.VertexAttribPointer(gl.Attrib{0}, 3, gl.FLOAT, false, 0, 0)

  // create empty buffer
  vbo = v.gl.CreateBuffer()
  // bind the new buffer
  v.gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
  // set the data (4 is sizeof float)
  v.gl.BufferData(gl.ARRAY_BUFFER, float_array_to_byte_array(uvs), gl.STATIC_DRAW)
  // bind vbo to shader
  // v.gl.VertexAttribPointer(gl.Attrib{1}, 2, gl.FLOAT, false, 0, 0)

  // create empty buffer
  vbo = v.gl.CreateBuffer()
  // bind the new buffer
  v.gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
  // set the data (4 is sizeof float)
  v.gl.BufferData(gl.ARRAY_BUFFER, float_array_to_byte_array(normals), gl.STATIC_DRAW)
  // bind vbo to shader
  // v.gl.VertexAttribPointer(gl.Attrib{2}, 3, gl.FLOAT, false, 0, 0)
}

func int_array_to_byte_array(arr []int32) []byte {
  buf := make([]byte, len(arr) * 4)
  for i, v := range arr {
    buf[i*4 + 0] = byte(v >> 24)
    buf[i*4 + 1] = byte(v >> 16)
    buf[i*4 + 2] = byte(v >> 8)
    buf[i*4 + 3] = byte(v)
  }
  return buf
}

func float_array_to_byte_array(arr []float32) []byte {
  buf := make([]byte, len(arr) * 4)
  for i, v := range arr {
    n := math.Float32bits(v)
    buf[i*4 + 0] = byte(n >> 24)
    buf[i*4 + 1] = byte(n >> 16)
    buf[i*4 + 2] = byte(n >> 8)
    buf[i*4 + 3] = byte(n)
  }
  return buf
}
