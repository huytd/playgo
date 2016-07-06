package engine_test

import (
	"testing"

	"github.com/huytd/go-play/engine"
)

func TestEngineCodeGeneratorForSingleLine(t *testing.T) {
	code := "print(\"Hello\")"
	expect := "package main\n\nfunc main() {\nprint(\"Hello\")\n}"

	e := engine.Engine{}
	output := e.Gen(code)
	if output != expect {
		t.Fail()
	}
}

func TestEngineCodeGeneratorForMultipleLines(t *testing.T) {
	code := "a := 1\n\tb := 2\nc := a + b\nprint(c)\n\n"
	expect := "package main\n\nfunc main() {\na := 1\n\tb := 2\nc := a + b\nprint(c)\n\n\n}"

	e := engine.Engine{}
	output := e.Gen(code)
	if output != expect {
		t.Fail()
	}
}

func TestEngineCodeGeneratorWithImport(t *testing.T) {
	//
	code := `import "strconv"

s := "5"
n, _ := strconv.Atoi(s)
print(n)`

	//
	expect := `package main
import "strconv"
func main() {

s := "5"
n, _ := strconv.Atoi(s)
print(n)
}`

	e := engine.Engine{}
	output := e.Gen(code)
	if output != expect {
		t.Logf("Expected:\n%s", expect)
		t.Logf("Got:\n%s", output)
		t.Fail()
	}
}
