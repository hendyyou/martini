package martini

import (
	"net/http"
	"os"
	"path/filepath"
)

func Static(path string) Handler {
	return func(res http.ResponseWriter, req *http.Request) {
		file := filepath.Join(path, req.URL.Path)
		info, err := os.Stat(file)
		if err == nil && !info.IsDir() {
			http.ServeFile(res, req, file)
		}
	}
}