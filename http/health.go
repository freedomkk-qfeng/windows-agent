package http

import (
	"net/http"

	"github.com/freedomkk-qfeng/windows-agent/g"
)

func configHealthRoutes() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(g.VERSION))
	})
}
