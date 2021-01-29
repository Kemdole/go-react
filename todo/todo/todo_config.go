package todo

import (
	"go-react/common"
	"net/http"
)

type Dependancies struct {
	Repository RepositoryIface
}

func PublicRouteConfigs(svc *Service) *common.RouteConfig {
	return &common.RouteConfig{
		RootPath: "/todos",
		Routes: []common.Route{
			{
				Endpoint: "",
				Method:   http.MethodPost,
				Handler:  createTodoHandler(svc),
			},
			{
				Endpoint: "",
				Method:   http.MethodGet,
				Handler:  getTodoListHandler(svc),
			},
			{
				Endpoint: "/:id",
				Method:   http.MethodGet,
				Handler:  getTodoHandler(svc),
			},
			{
				Endpoint: "/:id",
				Method:   http.MethodPut,
				Handler:  updateTodoHandler(svc),
			},
			{
				Endpoint: "/:id",
				Method:   http.MethodDelete,
				Handler:  deleteTodoHandler(svc),
			},
		},
	}
}

func PrivateRouteConfigs(svc *Service) *common.RouteConfig {
	return &common.RouteConfig{
		RootPath: "/todo",
		Routes:   nil,
	}
}
