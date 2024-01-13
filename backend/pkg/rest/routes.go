package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type routes struct {
	method  string
	path    string
	handler echo.HandlerFunc
}

func (r *Rest) loadRoutes(e *echo.Echo) {
	api := []routes{
		{"get", "v1/health", r.health},
		{"post", "v1/task", r.createTask},
		{"get", "v1/task/:id", r.getTask},
		{"get", "v1/task/owner/:owner", r.getAllTask},
		{"delete", "v1/task/:id", r.deleteTask},
		{"put", "v1/task/:id/done/:done", r.markDone},
	}
	for _, v := range api {
		switch v.method {
		case "get":
			e.GET(v.path, v.handler)
		case "post":
			e.POST(v.path, v.handler)
		case "put":
			e.PUT(v.path, v.handler)
		case "delete":
			e.DELETE(v.path, v.handler)
		default:
			log.Debug("skipping route: ", v.path)
			continue
		}
		log.Debug("route loaded: ", v.path)
	}

}
