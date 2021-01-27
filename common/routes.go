package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouteConfig struct {
	RootPath string
	Routes   []Route
	Inits    []InitFunc
}

type InitFunc func()

type Route struct {
	Endpoint   string
	Method     string
	Policy     string
	Version    string
	Handler    gin.HandlerFunc
	MiddleWare []gin.HandlerFunc
}

func RegisterRoutes(r *gin.Engine, c RouteConfig) {
	g := r.Group(c.RootPath)
	for i := range c.Routes {
		var ep string
		if c.Routes[i].Version != "" {
			ep = c.Routes[i].Version + c.Routes[i].Endpoint
		} else {
			ep = c.Routes[i].Endpoint
		}

		switch c.Routes[i].Method {
		case http.MethodGet:
			g.GET(ep, c.Routes[i].Handler)
		case http.MethodPost:
			g.POST(ep, c.Routes[i].Handler)
		case http.MethodPut:
			g.PUT(ep, c.Routes[i].Handler)
		case http.MethodDelete:
			g.DELETE(ep, c.Routes[i].Handler)
		case http.MethodHead:
			g.HEAD(ep, c.Routes[i].Handler)
		case http.MethodPatch:
			g.PATCH(ep, c.Routes[i].Handler)
		}

	}
}
