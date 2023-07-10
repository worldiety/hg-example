module github.com/worldiety/hg-example

go 1.20

require (
	github.com/go-chi/chi/v5 v5.0.8
	github.com/laher/mergefs v0.1.1
	github.com/vearutop/statigz v1.3.0
	github.com/worldiety/hg v0.0.0-20230710152952-dca84dfb09bc
	golang.org/x/exp v0.0.0-20230626212559-97b1e661b5df
)

//replace github.com/worldiety/hg => ../hg.git
