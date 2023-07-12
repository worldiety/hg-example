module github.com/worldiety/hg-example

go 1.20

require (
	github.com/go-chi/chi/v5 v5.0.8
	github.com/laher/mergefs v0.1.1
	github.com/vearutop/statigz v1.3.0
	github.com/worldiety/hg v0.0.0-20230712155952-4faf6355bff6
	golang.org/x/exp v0.0.0-20230711153332-06a737ee72cb
)

//replace github.com/worldiety/hg => ../hg.git
