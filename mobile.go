// +build android ios

package gl

import (
  "golang.org/x/mobile/app"
  "golang.org/x/mobile/event/lifecycle"
  "golang.org/x/mobile/event/paint"
  "golang.org/x/mobile/event/size"
  "golang.org/x/mobile/event/touch"
  "golang.org/x/mobile/gl"

  "github.com/macmv/simple-gl/core"
)

func Main() {
  app.Main(func(a app.App) {
    var glctx gl.Context
    for e := range a.Events() {
      switch e := a.Filter(e).(type) {
      case lifecycle.Event:
        switch e.Crosses(lifecycle.StageVisible) {
        case lifecycle.CrossOn:
          glctx, _ = e.DrawContext.(gl.Context)
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
        a.Publish()
        // Drive the animation by preparing to paint the next frame
        // after this one is shown.
        a.Send(paint.Event{})
      case touch.Event:
        core.Touch(e)
      }
    }
  })
}
