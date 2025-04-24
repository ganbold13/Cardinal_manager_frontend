package dist

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed files/*
var embeddedFiles embed.FS

func New() http.FileSystem {
	fsys, err := fs.Sub(embeddedFiles, "files/manager")
	if err != nil {
		panic(err)
	}
	return WithSPA(http.FS(fsys))
}
