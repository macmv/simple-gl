package gl

import (
  "reflect"
)

var (
  events map[int]reflect.Value
)

func On(name int, handle interface{}) {
  v := reflect.ValueOf(handle)
  if v.Kind() != reflect.Func {
    panic("Please pass a function into relfect.On")
  }
  if events == nil {
    events = make(map[int]reflect.Value)
  }
  events[name] = v
}

func run(name int, values... interface{}) {
  handle, ok := events[name]
  if !ok {
    return
  }
  rvalues := []reflect.Value{}
  for _, v := range values {
    rvalues = append(rvalues, reflect.ValueOf(v))
  }
  handle.Call(rvalues)
}
