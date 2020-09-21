// !build android ios

package gl

import (
  "fmt"

  "golang.org/x/mobile/gl"
  "golang.org/x/exp/shiny/driver"
  "golang.org/x/exp/shiny/screen"
  "golang.org/x/mobile/event/key"
  "golang.org/x/mobile/event/size"
  "golang.org/x/mobile/event/paint"
  "golang.org/x/mobile/event/touch"
  "golang.org/x/mobile/event/lifecycle"

  "github.com/macmv/simple-gl/core"
  "github.com/macmv/simple-gl/event"
  "github.com/macmv/simple-gl/desktop"
)

func Main() {
  driver.Main(func(s screen.Screen) {
    w, err := s.NewWindow(nil)
    if err != nil {
      panic(err)
      return
    }
    defer w.Release()

    var c core.Core = desktop.NewCore()

    var glctx gl.Context
    for {
      switch e := w.NextEvent().(type) {
      case lifecycle.Event:
        run(DRAW, event.Draw{}, c)
        if e.To == lifecycle.StageDead {
          return
        }

        switch e.Crosses(lifecycle.StageVisible) {
        case lifecycle.CrossOn:
          core.Start(glctx)
        case lifecycle.CrossOff:
          core.Stop()
        }
      case size.Event:
        core.Resize(e)
      case paint.Event:
        if glctx == nil || e.External {
          // As we are actively painting as fast as
          // we can (usually 60 FPS), skip any paint
          // events sent by the system.
          continue
        }

        core.Paint()
        w.Publish()
        // Drive the animation by preparing to paint the next frame
        // after this one is shown.
        w.Send(paint.Event{})
      case touch.Event:
        core.Touch(e)
      case key.Event:
        core.Key(e)
        fmt.Println("Keyboard!", e)
      }
    }
  })
}

