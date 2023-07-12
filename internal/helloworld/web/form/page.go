package form

import (
	"embed"
	"github.com/worldiety/hg"
	"github.com/worldiety/hg-example/internal/helloworld"
	"io/fs"
	"net/http"
	"strconv"
)

//go:embed *.gohtml
var templateFiles embed.FS

type PageState struct {
	Title   string
	DevMode bool
	Form    FormModel
}

type CheckEvent struct {
	Firstname string
	Age       string
}

func (evt CheckEvent) ToDomain() (helloworld.Person, helloworld.PersonValidationError) {
	var p helloworld.Person
	age, err := strconv.Atoi(evt.Age)
	if err != nil {
		// this is discussable, how harsh our reaction is.
		// actually, this does not make much sense, because we could tell the form, that we expect an int,
		// we could fail with something technically, instead of this unsuited domain error
		return p, helloworld.InvalidAgeError(-1)
	}

	p.Age = age
	p.Firstname = evt.Firstname

	return p, helloworld.PersonCheck(p)
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
				file.RenameTplStmts("page", "Form")
			}),
		),

		hg.OnRequest(
			func(r *http.Request, model PageState) PageState {
				model.Title = "greetings from form"
				model.DevMode = true
				return model
			},
		),

		hg.Case("Check", func(model PageState, msg CheckEvent) PageState {
			model.Form.AgeError = ""
			model.Form.FirstnameError = ""
			person, err := msg.ToDomain()

			// what is not very nice, but still better than average, is that the user has to solve one error after
			// another. Either way, the input-data should not be lost.
			switch err.(type) {
			case helloworld.UnexpectedNameError:
				model.Form.FirstnameError = "this is not the expected name, try 'Torben'"
			case helloworld.InvalidAgeError:
				model.Form.AgeError = "Age is invalid. Must be a number between 18 and 120"
			}

			model.Form.Age = person.Age
			model.Form.Firstname = person.Firstname

			return model
		}),
	)
}
