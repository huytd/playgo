package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/huytd/playgo/engine"
	"github.com/huytd/playgo/web"
)

var (
	mode = flag.String("mode", "cli", "running mode of the playground (cli or web)")
)

func main() {
	flag.Parse()
	switch *mode {
	case "web":
		web.Start()
	case "cli":
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("Can't read standard input: %v", err)
		}
		code := string(bytes)
		if code == "" {
			log.Fatal("No code to run!")
		}
		e := engine.Engine{}
		output, err := e.Run(code)
		if err != nil {
			log.Fatalf("Error executing code: %v", err)
		}
		fmt.Print(output)
	default:
		log.Fatal("Oops! There are only 2 modes supported: cli and web")
	}
}
