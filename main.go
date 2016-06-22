package main

import "github.com/huytd/go-play/engine"

func main() {
	e := engine.Engine{}
	e.Run(`print("Hello World")
  a := 1
  b := 2
  print("Result: ", a + b)`)
}
