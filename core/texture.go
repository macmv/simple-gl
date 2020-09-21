package core

type Texture interface {
  Data(data []uint8)
  Bind()
}
