package routes

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/maker-space-experimenta/printer-kiosk/util"
)

type SpaHandler struct {
	config     util.Config
	staticPath string
	indexPath  string
}

func NewSpaHandler(config util.Config, staticPath string, indexPath string) *SpaHandler {
	return &SpaHandler{
		staticPath: staticPath,
		indexPath:  indexPath,
		config:     config,
	}
}

func (h *SpaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
