package core

type Core interface {
  NewVAO() VAO
  NewShaderGeo(geometry_path, vertex_path, fragment_path string) (Shader, error)
  NewTexture3DFromData(width, height, depth int, data []uint8) Texture
  Window() Window
}

