package webserver

import (
	"os/exec"

	"github.com/go-martini/martini"
	"github.com/johnsudaar/htagrunner/config"
)

func Start() {
	m := martini.Classic()
	m.Get("/:app/:branch/:version", func(params martini.Params) string {
		app := "runner"
		cmd := exec.Command(app, params["app"], params["branch"], params["version"])
		stdout, err := cmd.Output()
		if err != nil {
			return string(stdout) + err.Error()
		} else {
			return string(stdout)
		}
	})
	m.RunOnAddr(config.E["PORT"])
}
