package main

import "github.com/huytd/go-play/engine"

func main() {
	e := engine.Engine{}
	e.Run(`println("Hello World")
  a := 1
  b := 2
  println("Result: ", a + b)`)
}
