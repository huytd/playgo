package web

import (
	"github.com/huytd/go-play/engine"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
)

func runCode(w http.ResponseWriter, r *http.Request) {
	e := engine.Engine{}
	output := e.Run(r.FormValue("code"))
	io.WriteString(w, string(output))
}

func Start() {
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":3000"
	}

	println("Web IDE started at http://localhost", port)

	_, filename, _, _ := runtime.Caller(1)
	fs := http.FileServer(http.Dir(path.Join(path.Dir(filename), "www")))
	http.Handle("/", fs)
	http.HandleFunc("/api/run", runCode)
	http.ListenAndServe(port, nil)
}
