package main

import (
	"github.com/spf13/viper"
	"go-react/service"
	"go-react/todo/todo"
)

func main() {
	service.StartServer(
		[]*service.RouteConfig{
			todo.PublicRouteConfigs,
		},
		viper.GetInt("server.port"),
	)
	service.StartServer(
		[]*service.RouteConfig{
			todo.PrivateRouteConfigs,
		},
		viper.GetInt("server.privatePort"),
	)
}
