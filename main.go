package main

import (
	"flag"
	"github.com/huytd/go-play/engine"
	"github.com/huytd/go-play/web"
	"io/ioutil"
	"os"
)

var (
	modeFlag = flag.String("mode", "cli", "running mode of the playground")
)

func main() {
	flag.Parse()

	mode := *modeFlag
	if mode == "cli" {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			println("No code to run!")
		} else {
			code := string(bytes)
			if code != "" {
				e := engine.Engine{}
				output := e.Run(code)
				println(output)
			} else {
				println("No code to run!")
			}
		}
	} else if mode == "web" {
		web.Start()
	} else {
		println("Oops! There are only 2 modes supported: cli and web")
	}
}
