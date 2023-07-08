package helloworld_srv

import "github.com/worldiety/hg-example/internal/app/helloworld-srv/web"

func Run(host string) error {
	return web.Serve(host)
}
