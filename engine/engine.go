package engine

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

type Engine struct {
}

// Capture execute an anonymous function and capture the os.Stdout
func (e *Engine) Capture(fn func(*os.File, string) error, params string) (string, error) {
	originalStdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outChannel := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outChannel <- buf.String()
	}()

	err := fn(w, params)

	w.Close()
	os.Stdout = originalStdOut
	output := <-outChannel

	return output, err
}

// Format run gofmt to format the code
func (e *Engine) Format(input string) (string, error) {
	out, err := e.Capture(func(w *os.File, code string) error {
		path, _ := e.Save(e.Gen(code))
		defer e.CleanUp(path)

		os.Chdir(path)

		cmdName := "gofmt"
		cmdArgs := []string{"main.go"}
		cmd := exec.Command(cmdName, cmdArgs...)
		cmd.Stdout = w
		cmd.Stderr = w

		return cmd.Run()
	}, input)
	return out, err
}

// Gen is a function to preprocess the code before execute it
func (e *Engine) Gen(input string) string {
	imports := []string{}
	statements := []string{}

	// Check code starts with "package main"
	if m, _ := regexp.MatchString("^\\s*package\\s+main\\s*", input); m {
		// No preprocessing required
		return input
	}

	// Append package and all import
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

// Save function write the input code to a temporary folder for processing
func (e *Engine) Save(code string) (string, string) {
	dir, err := ioutil.TempDir("", "playgo-engine-temp")
	if err != nil {
		log.Fatal(err)
	}
	tmpfn := filepath.Join(dir, "main.go")
	if err := ioutil.WriteFile(tmpfn, []byte(code), 0666); err != nil {
		log.Fatal(err)
	}
	return dir, tmpfn
}

// CleanUp will delete the temporary folder after executing
func (e *Engine) CleanUp(dir string) {
	os.RemoveAll(dir)
}

// Run function execute the code and return the output
func (e *Engine) Run(input string) (string, error) {
	out, err := e.Capture(func(w *os.File, code string) error {
		path, _ := e.Save(e.Gen(code))
		defer e.CleanUp(path)

		os.Chdir(path)

		cmdName := "go"
		cmdArgs := []string{"run", "main.go"}
		cmd := exec.Command(cmdName, cmdArgs...)
		cmd.Stdout = w
		cmd.Stderr = w

		return cmd.Run()
	}, input)
	return out, err
}
