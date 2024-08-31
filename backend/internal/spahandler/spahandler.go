package spahandler

import (
	"net/http"

	"github.com/retrokyo/basedchat/internal/frontend"
)

type SpaHandler struct {
	StaticPath string
	IndexPath  string
}

func (handler SpaHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	if !frontend.EmbeddedFS.Exists(handler.StaticPath, req.URL.Path) {
		http.FileServer(frontend.FallbackFS).ServeHTTP(writer, req)
		return
	}

	http.FileServer(frontend.EmbeddedFS).ServeHTTP(writer, req)
}
