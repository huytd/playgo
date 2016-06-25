package engine_test

import (
	"github.com/huytd/go-play/engine"
	"testing"
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
	code := "import \"strconv\"\n\ns := \"5\"\nn, _ := strconv.Atoi(s)\nprint(n)"
	expect := "package main\n\nimport \"strconv\"\n\nfunc main() {\ns := \"5\"\nn, _ := strconv.Atoi(s)\nprint(n)\n}"

	e := engine.Engine{}
	output := e.Gen(code)
	if output != expect {
		t.Fail()
	}
}
