// +build !android ios

package gl

import (
  "fmt"

  "golang.org/x/mobile/gl"
  "golang.org/x/exp/shiny/screen"
  "golang.org/x/mobile/event/key"
  "golang.org/x/mobile/event/size"
  "golang.org/x/mobile/event/paint"
  "golang.org/x/mobile/event/touch"
  "golang.org/x/mobile/event/lifecycle"
  "golang.org/x/exp/shiny/driver/gldriver"

  "github.com/macmv/simple-gl/core"
  "github.com/macmv/simple-gl/event"
  "github.com/macmv/simple-gl/desktop"
)

func Main() {
  gldriver.Main(func(s screen.Screen) {
    w, err := s.NewWindow(nil)
    if err != nil {
      panic(err)
      return
    }
    defer w.Release()

    c := desktop.NewCore()

    var glctx gl.Context3
    for {
      switch e := w.NextEvent().(type) {
      case lifecycle.Event:
        if e.To == lifecycle.StageDead {
          return
        }

        switch e.Crosses(lifecycle.StageVisible) {
        case lifecycle.CrossOn:
          glctx, _ = e.DrawContext.(gl.Context3)
          c.SetWindow(glctx, w)
          run(START, core.Core(c))
        case lifecycle.CrossOff:
          core.Stop()
        }
      case size.Event:
        core.Resize(e)
      case paint.Event:
        run(DRAW, event.Draw{}, c)
        core.Paint()
        w.Publish()
      case touch.Event:
        core.Touch(e)
      case key.Event:
        core.Key(e)
        fmt.Println("Keyboard!", e)
      }
    }
  })
}

