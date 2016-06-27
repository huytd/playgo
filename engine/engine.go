package engine

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Engine struct {
}

func (e *Engine) Gen(input string) string {
	imports := []string{}
	statements := []string{}

	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "import ") {
			imports = append(imports, line)
		} else {
			statements = append(statements, line)
		}
	}

	generated := "package main\n"
	generated = generated + strings.Join(imports, "\n") + "\n"
	generated = generated + "func main() {\n" + strings.Join(statements, "\n") + "\n}"
	return generated
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

func (e *Engine) Run(code string) (string, error) {
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

	err := cmd.Run()

	w.Close()
	os.Stdout = originalStdOut
	output := <-outChannel

	return output, err
}
