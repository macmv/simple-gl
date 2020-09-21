package desktop

import (
  "github.com/macmv/simple-gl/core"
)

type Core struct {
  window *Window
}

func NewCore() *Core {
  c := Core{}
  return &c
}

func (c *Core) Window() core.Window {
  return c.window
}
