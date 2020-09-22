package desktop

import (
  "os"
  "fmt"
  "image"
  "image/draw"
  _ "image/png"

  "github.com/macmv/simple-gl/core"

  "golang.org/x/mobile/gl"
)

type Texture struct {
  dimension int
  tex gl.Texture
  gl gl.Context3
  width, height, depth int
}

func (c *Core) NewTexture2DFromData(width, height int, data []byte) core.Texture {
  tex := c.gl.CreateTexture()
  c.gl.BindTexture(gl.TEXTURE_2D, tex)
  c.gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
  c.gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
  c.gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
  c.gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
  c.gl.TexImage2D(
    gl.TEXTURE_2D,
    0,
    gl.RGBA,
    width,
    height,
    gl.RGBA,
    gl.UNSIGNED_BYTE,
    data,
  )

  t := Texture{}
  t.gl = c.gl
  t.tex = tex
  t.width = width
  t.height = height
  t.dimension = 2
  return &t
}

func (c *Core) NewTexture2DFromFile(path string) (core.Texture, error) {
  img_file, err := os.Open(path)
  if err != nil {
    return nil, fmt.Errorf("Texture %q not found on disk: %v", path, err)
  }
  img, _, err := image.Decode(img_file)
  if err != nil {
    return nil, err
  }

  rgba := image.NewRGBA(img.Bounds())
  if rgba.Stride != rgba.Rect.Size().X * 4 {
    return nil, fmt.Errorf("Unsupported stride")
  }
  draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

  return c.NewTexture2DFromData(rgba.Rect.Size().X, rgba.Rect.Size().Y, rgba.Pix), nil
}

func (t *Texture) Data(data []uint8) {
  t.Bind()
  if t.dimension == 2 {
    t.gl.TexImage2D(
      gl.TEXTURE_2D,
      0,
      gl.RGBA,
      t.width,
      t.height,
      gl.RGBA,
      gl.UNSIGNED_BYTE,
      data,
    )
  } else if t.dimension == 3 {
    t.gl.TexImage2D(
      gl.TEXTURE_3D,
      0,
      gl.RGBA,
      t.width,
      t.height,
      gl.RGBA,
      gl.UNSIGNED_BYTE,
      data,
    )
  } else {
    panic("Unsuported dimension! (should be 2 or 3 for 2d or 3d textures)")
  }
}

func (t *Texture) Bind() {
  if t.dimension == 2 {
    t.gl.BindTexture(gl.TEXTURE_2D, t.tex)
  } else if t.dimension == 3 {
    t.gl.BindTexture(gl.TEXTURE_3D, t.tex)
  } else {
    panic("Unsuported dimension! (should be 2 or 3 for 2d or 3d textures)")
  }
}
