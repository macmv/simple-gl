package core

import (
)

type VAO interface {
  SetData([]int32, []float32, []float32, []float32)
  Length() int
  Bind()
  Unbind()
}
