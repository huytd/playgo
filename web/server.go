package web

import (
	"io"
	"log"
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

func formatCode(w http.ResponseWriter, r *http.Request) {
	e := engine.Engine{}
	output, err := e.Format(r.FormValue("code"))
	if err != nil {
		http.Error(w, output, 400)
		return
	}
	io.WriteString(w, string(output))
}

const defaultPort = "3000"

// Start launches an HTTP server and attaches handlers.
func Start() {
	port := envString("PLAYGO_PORT", defaultPort)
	log.Printf("Starting web IDE at http://localhost:%s", port)

	_, filename, _, _ := runtime.Caller(1)
	fs := http.FileServer(http.Dir(path.Join(filepath.Dir(filename), "www")))

	http.Handle("/", fs)
	http.HandleFunc("/api/run", runCode)
	http.HandleFunc("/api/format", formatCode)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Printf("Can't start HTTP listener: %v", err)
	}
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
