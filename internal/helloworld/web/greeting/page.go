package greeting

import (
	"embed"
	"github.com/worldiety/hg"
	"github.com/worldiety/hg-example/internal/helloworld"
	"io/fs"
	"net/http"
)

//go:embed *.gohtml
var templateFiles embed.FS

type PageState struct {
	Title   string
	Greets  string
	DevMode bool `json:"-"`
	Count   int
}

type AddEvent int

func Page(tplFiles ...fs.FS) http.HandlerFunc {
	return hg.Handler(
		hg.MustParse[PageState](
			hg.FS(tplFiles...),
			hg.FS(templateFiles),
			hg.Execute("index"),
			// one possible way to handle "unknown" forward template references (see index.gohtml):
			// simple and efficient, just rename the generic page placeholder with our component name
			hg.TplPreProcEach(func(file *hg.TplFile) {
				file.RenameTplStmts("page", "greeting")
			}),
		),

		// onRequest is always called for anything that matches the registered route.
		hg.OnRequest(
			func(r *http.Request, model PageState) PageState {
				model.Title = "greetings from hg"
				model.DevMode = true
				name := "Torben"
				if optName := r.URL.Query().Get("name"); optName != "" {
					name = optName
				}

				// call into our business (domain) layer
				model.Greets = helloworld.SayHello(name)
				return model
			},
		),

		hg.Update(hg.Case("add", func(model PageState, msg AddEvent) PageState {
			model.Count += int(msg)
			return model
		})),
	)
}
