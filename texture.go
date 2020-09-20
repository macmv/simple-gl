package gl

import (
  "os"
  "fmt"
  "image"
  "image/draw"
  _ "image/png"

  "github.com/go-gl/gl/v4.1-core/gl"
)

type Texture struct {
  dimension int
  id uint32
  width, height, depth int
}

func NewTexture3DFromData(width, height, depth int, data []uint8) *Texture {
  var id uint32
  gl.GenTextures(1, &id)
  gl.BindTexture(gl.TEXTURE_3D, id)
  gl.TexParameteri(gl.TEXTURE_3D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
  gl.TexParameteri(gl.TEXTURE_3D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
  gl.TexParameteri(gl.TEXTURE_3D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
  gl.TexParameteri(gl.TEXTURE_3D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
  gl.TexImage3D(
    gl.TEXTURE_3D,
    0,
    gl.RGBA,
    int32(width),
    int32(height),
    int32(depth),
    0,
    gl.RGBA,
    gl.UNSIGNED_BYTE,
    gl.Ptr(data),
  )

  t := Texture{}
  t.id = id
  t.width = width
  t.height = height
  t.depth = depth
  t.dimension = 3
  return &t
}

func NewTexture2DFromData(width, height int, data []uint8) *Texture {
  var id uint32
  gl.GenTextures(1, &id)
  gl.BindTexture(gl.TEXTURE_2D, id)
  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
  gl.TexImage2D(
    gl.TEXTURE_2D,
    0,
    gl.RGBA,
    int32(width),
    int32(height),
    0,
    gl.RGBA,
    gl.UNSIGNED_BYTE,
    gl.Ptr(data),
  )

  t := Texture{}
  t.id = id
  t.width = width
  t.height = height
  t.dimension = 2
  return &t
}

func NewTexture2DFromFile(path string) (*Texture, error) {
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

  return NewTexture2DFromData(rgba.Rect.Size().X, rgba.Rect.Size().Y, rgba.Pix), nil
}

func (t *Texture) Data(data []uint8) {
  t.Bind()
  if t.dimension == 2 {
    gl.TexImage2D(
      gl.TEXTURE_2D,
      0,
      gl.RGBA,
      int32(t.width),
      int32(t.height),
      0,
      gl.RGBA,
      gl.UNSIGNED_BYTE,
      gl.Ptr(data),
    )
  } else if t.dimension == 3 {
    gl.TexImage3D(
      gl.TEXTURE_3D,
      0,
      gl.RGBA,
      int32(t.width),
      int32(t.height),
      int32(t.depth),
      0,
      gl.RGBA,
      gl.UNSIGNED_BYTE,
      gl.Ptr(data),
    )
  } else {
    panic("Unsuported dimension! (should be 2 or 3 for 2d or 3d textures)")
  }
}

func (t *Texture) Bind() {
  if t.dimension == 2 {
    gl.BindTexture(gl.TEXTURE_2D, t.id)
  } else if t.dimension == 3 {
    gl.BindTexture(gl.TEXTURE_3D, t.id)
  } else {
    panic("Unsuported dimension! (should be 2 or 3 for 2d or 3d textures)")
  }
}
