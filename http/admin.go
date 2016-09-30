package http

import (
	"net/http"

	"github.com/freedomkk-qfeng/windows-agent/g"
	"github.com/toolkits/file"
)

func configAdminRoutes() {

	http.HandleFunc("/workdir", func(w http.ResponseWriter, r *http.Request) {
		RenderDataJson(w, file.SelfDir())
	})

	http.HandleFunc("/ips", func(w http.ResponseWriter, r *http.Request) {
		RenderDataJson(w, g.TrustableIps())
	})
}
