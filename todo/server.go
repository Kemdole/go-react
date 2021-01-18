package main

import (
	"go-react/service"
	"go-react/todo/todo"
)

func main() {
	service.StartService(service.ServerConfig{
		ServiceName:       "todo",
		PublicRouteConfig: todo.PublicRouteConfigs,
		PrivateRoutConfig: todo.PrivateRouteConfigs,
	})
}
