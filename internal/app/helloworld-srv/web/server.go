package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/worldiety/hg-example/internal/app/helloworld-srv/web/index"
	"github.com/worldiety/hg-example/internal/helloworld/web/greeting"
	"net/http"
)

func Serve(host string) error {
	router := chi.NewRouter()
	greeting.Route(router, index.TemplateFiles)

	return http.ListenAndServe(host, router)
}
