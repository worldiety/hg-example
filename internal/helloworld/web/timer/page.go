package timer

import (
	"embed"
	"github.com/worldiety/hg"
	"io/fs"
	"net/http"
	"time"
)

//go:embed *.gohtml
var templateFiles embed.FS

type PageState struct {
	Title   string
	DevMode bool
	Time    string
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
				file.RenameTplStmts("page", "Timer")
			}),
		),

		hg.OnRequest(
			func(r *http.Request, model PageState) PageState {
				model.Title = "Timer every second example"
				model.DevMode = true
				model.Time = time.Now().String()
				return model
			},
		),
	)
}
