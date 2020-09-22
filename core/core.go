package core

import (
  "golang.org/x/mobile/gl"
)

type Core interface {
  NewVAO() VAO
  NewTexture2DFromData(width, height int, data []byte) Texture
  Window() Window
  Gl() gl.Context3
}

