package main

import (
	"github.com/spf13/viper"
	"go-react/common"
	"go-react/todo/todo"
)

func main() {
	done := make(chan bool)
	defer close(done)

	common.ReadConfigs("monolith")

	todoCtr := todo.NewController()

	go common.StartServer(
		[]*common.RouteConfig{
			todo.PublicRouteConfigs(todoCtr),
		},
		viper.GetInt("server.port"),
		done,
	)
	go common.StartServer(
		[]*common.RouteConfig{
			todo.PrivateRouteConfigs(todoCtr),
		},
		viper.GetInt("server.privatePort"),
		done,
	)

	<-done
	<-done
}
