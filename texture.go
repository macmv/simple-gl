package gl

import (
  "github.com/go-gl/gl/v4.1-core/gl"
)

type Texture struct {
  id uint32
}

func (t *Texture) Bind() {
  gl.BindTexture(gl.TEXTURE_2D, t.id)
}
