package frontend

import (
	"fmt"
	"io/fs"
	"net/http"
	"strings"
)

type StaticFileSystem struct {
	http.FileSystem
}

var _ http.FileSystem = (*StaticFileSystem)(nil)

func newStaticFileSystem() *StaticFileSystem {
	sub, err := fs.Sub(Content, "build")

	if err != nil {
		panic(err)
	}

	return &StaticFileSystem{
		FileSystem: http.FS(sub),
	}
}

func (sfs *StaticFileSystem) Exists(prefix string, path string) bool {
	buildpath := fmt.Sprintf("build%s", path)

	if strings.HasSuffix(path, "/") {
		_, err := Content.ReadDir(strings.TrimSuffix(buildpath, "/"))
		return err == nil
	}

	file, err := Content.Open(buildpath)
	if file != nil {
		_ = file.Close()
	}

	return err == nil
}

type FallbackFileSystem struct {
	staticFileSystem *StaticFileSystem
}

var _ http.FileSystem = (*FallbackFileSystem)(nil)

func (fallback *FallbackFileSystem) Open(path string) (http.File, error) {
	return fallback.staticFileSystem.Open("/index.html")
}

var EmbeddedFS *StaticFileSystem = newStaticFileSystem()
var FallbackFS *FallbackFileSystem = &FallbackFileSystem{
	staticFileSystem: EmbeddedFS,
}
