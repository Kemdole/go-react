package todo

import "go-react/service"

var PublicRouteConfigs = &service.RouteConfig{
	RootPath: "/todo",
	Routes:   nil,
	Inits:    []service.InitFunc{},
}

var PrivateRouteConfigs = &service.RouteConfig{
	RootPath: "/todo",
	Routes:   nil,
	Inits:    []service.InitFunc{},
}
