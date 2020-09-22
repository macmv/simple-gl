package desktop

import (
  "github.com/macmv/simple-gl/core"

  "golang.org/x/exp/shiny/screen"
  "golang.org/x/mobile/gl"
)

type Core struct {
  window *Window
  gl gl.Context3
}

func NewCore() *Core {
  c := Core{}
  return &c
}

func (c *Core) SetWindow(glctx gl.Context3, w screen.Window) {
  c.gl = glctx
  c.window = new_window(glctx, w)
}

func (c *Core) Window() core.Window {
  return c.window
}

func (c *Core) Gl() gl.Context3 {
  return c.gl
}
