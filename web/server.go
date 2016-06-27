package web

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/huytd/go-play/engine"
)

func runCode(w http.ResponseWriter, r *http.Request) {
	e := engine.Engine{}
	output, err := e.Run(r.FormValue("code"))

	if err != nil {
		http.Error(w, output, 400)
		return
	}

	io.WriteString(w, string(output))
}

func Start() {
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":3000"
	}

	println("Web IDE started at http://localhost", port)

	_, filename, _, _ := runtime.Caller(1)
	fs := http.FileServer(http.Dir(path.Join(filepath.Dir(filename), "www")))
	http.Handle("/", fs)
	http.HandleFunc("/api/run", runCode)
	http.ListenAndServe(port, nil)
}
