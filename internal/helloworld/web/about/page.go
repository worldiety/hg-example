package about

import (
	"embed"
	"github.com/worldiety/hg"
	"io/fs"
	"net/http"
)

//go:embed *.gohtml
var templateFiles embed.FS

type PageState struct {
	Title   string
	DevMode bool
}

func Page(tplFiles ...fs.FS) http.HandlerFunc {
	return hg.Handler(
		hg.MustParse[PageState](
			hg.FS(tplFiles...),
			hg.FS(templateFiles),
			hg.Execute("index"),
			// one possible way to handle "unknown" forward template references (see index.gohtml):
			// simple and efficient, just rename the generic page placeholder with our component name
			hg.TplPreProcEach(func(file *hg.TplFile) {
				file.RenameTplStmts("page", "About")
			}),
		),

		hg.OnRequest(
			func(r *http.Request, model PageState) PageState {
				model.Title = "greetings from about"
				model.DevMode = true
				return model
			},
		),
	)
}
