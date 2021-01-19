package todo

import (
	"go-react/common"
	"net/http"
)

func PublicRouteConfigs(svc *Service) *common.RouteConfig {
	return &common.RouteConfig{
		RootPath: "/todo",
		Routes: []common.Route{
			{
				Endpoint: "/:id",
				Method:   http.MethodGet,
				Handler:  getTodoHandler(svc),
			},
		},
		Inits: []common.InitFunc{
			initService(svc),
		},
	}
}

func PrivateRouteConfigs(svc *Service) *common.RouteConfig {
	return &common.RouteConfig{
		RootPath: "/todo",
		Routes:   nil,
		Inits:    []common.InitFunc{},
	}
}
