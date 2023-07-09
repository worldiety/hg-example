package helloworld

import "github.com/worldiety/hg-example/internal/app/helloworld/web"

func Run(host string) error {
	return web.Serve(host)
}
