package core

type Window interface {
  Use(shader Shader)
  Finish()

  Width() int
  Height() int
}
