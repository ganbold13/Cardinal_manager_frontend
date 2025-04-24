package dist

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed files/*
var embeddedFiles embed.FS

func New() http.FileSystem {
	// Use fs.Sub to get to the "files" subdir
	fsys, err := fs.Sub(embeddedFiles, "files")
	if err != nil {
		panic(err)
	}
	return WithSPA(http.FS(fsys))
}
