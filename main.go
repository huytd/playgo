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
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Printf("\ncli mode usage,\n  echo 'print(\"Hello, do some math, 1 + 1 = \", 1 + 1)' | %s \n    or\n  cat something.txt | %s\n", os.Args[0], os.Args[0])
	}
	flag.Parse()
	switch *mode {
	case "web":
		web.Start()
	case "cli":
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
			fmt.Println("Incorrect Usage\n")
			flag.Usage()
			os.Exit(1)
		}
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
