package main

import (
	"go-react/common"
	"go-react/todo/todo"
)

func main() {
	svc := todo.NewService()

	common.StartService(common.ServerConfig{
		ServiceName:       "todo",
		PublicRouteConfig: todo.PublicRouteConfigs(svc),
		PrivateRoutConfig: todo.PrivateRouteConfigs(svc),
	})
}
