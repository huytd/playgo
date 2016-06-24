package engine

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type Engine struct {
}

func (e *Engine) Gen(input string) string {
	return "package main\n\nfunc main() {\n" + input + "\n}"
}

func (e *Engine) Save(code string) (string, string) {
	dir, err := ioutil.TempDir("", "go-play-engine-temp")
	if err != nil {
		log.Fatal(err)
	}
	tmpfn := filepath.Join(dir, "main.go")
	if err := ioutil.WriteFile(tmpfn, []byte(code), 0666); err != nil {
		log.Fatal(err)
	}
	return dir, tmpfn
}

func (e *Engine) CleanUp(dir string) {
	os.RemoveAll(dir)
}

func (e *Engine) Run(code string) string {
	originalStdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outChannel := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outChannel <- buf.String()
	}()

	path, _ := e.Save(e.Gen(code))
	defer e.CleanUp(path)

	os.Chdir(path)

	cmdName := "go"
	cmdArgs := []string{"run", "main.go"}
	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Stdout = w
	cmd.Stderr = w

	cmd.Run()

	w.Close()
	os.Stdout = originalStdOut
	output := <-outChannel

	return output
}
