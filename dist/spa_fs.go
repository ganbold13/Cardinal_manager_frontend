// File: dist/spa_fs.go (optional name)
package dist

import (
	"net/http"
	"path"
)

// WithSPA wraps a FileSystem and returns index.html for missing files (for SPA routing).
func WithSPA(fs http.FileSystem) http.FileSystem {
	return &spaFS{fs}
}

type spaFS struct {
	fs http.FileSystem
}

func (s *spaFS) Open(name string) (http.File, error) {
	// Try to open the original file
	f, err := s.fs.Open(name)
	if err == nil {
		return f, nil
	}

	// If it's an asset (js, css, png, etc), just return the error
	if ext := path.Ext(name); ext != "" {
		return nil, err
	}

	// Otherwise, fallback to index.html
	return s.fs.Open("/index.html")
}
