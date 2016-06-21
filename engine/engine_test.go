package engine_test

import (
  "testing"
  "github.com/huytd/go-play/engine"
)

func TestEngineCodeGeneratorForSingleLine(t *testing.T) {
  code := "print(\"Hello\")"
  expect := "package main\n\nfunc main() {\n\tprint(\"Hello\")\n}"

  e := engine.Engine{}
  output := e.Gen(code)
  if output != expect {
    t.Fail()
  }
}

