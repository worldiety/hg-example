package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/laher/mergefs"
	"github.com/vearutop/statigz"
	"github.com/worldiety/hg"
	"github.com/worldiety/hg-example/internal/app/helloworld/web/index"
	"github.com/worldiety/hg-example/internal/helloworld/web/greeting"
	"net/http"
	"time"
)

func Serve(host string) error {
	router := chi.NewRouter()
	assets := statigz.FileServer(mergefs.Merge(hg.Assets, hg.Tailwind).(mergefs.MergedFS), statigz.EncodeOnInit)
	router.Handle("/assets/*", assets)
	router.Handle("/version/poll", hg.LongPollHandler(time.Second*30))
	greeting.Route(router, index.TemplateFiles)

	return http.ListenAndServe(host, router)
}
