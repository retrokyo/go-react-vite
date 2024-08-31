package spahandler

import (
	"net/http"
	"os"
	"path/filepath"
)

type SpaHandler struct {
	StaticPath string
	IndexPath  string
}

func (handler SpaHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	path := filepath.Join(handler.StaticPath, req.URL.Path)

	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {
		http.ServeFile(writer, req, filepath.Join(handler.StaticPath, handler.IndexPath))
		return
	}

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(handler.StaticPath)).ServeHTTP(writer, req)
}
