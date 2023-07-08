package greeting

import (
	"github.com/go-chi/chi/v5"
	"io/fs"
)

func Route(router chi.Router, tplFiles ...fs.FS) {
	router.HandleFunc("/", Page(tplFiles...)) // handler must respond to GET and POST requests
}
