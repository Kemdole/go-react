package common

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

type ServiceConfig struct {
	ServiceName       string
	PublicRouteConfig *RouteConfig
	PrivateRoutConfig *RouteConfig
	Inits             []InitFunc
	Events            []Event
}

type InitFunc func()
type Event interface{} //temp

func StartService(sc ServiceConfig) {
	ReadConfigs(sc.ServiceName)

	done := make(chan bool)
	defer close(done)

	publicPort := viper.GetInt("server.port")
	privatePort := viper.GetInt("server.privatePort")

	if sc.PublicRouteConfig == nil {
		log.Fatal("PublicRouteConfig")
	}
	go StartServer([]*RouteConfig{sc.PublicRouteConfig}, publicPort, done)

	if sc.PrivateRoutConfig != nil {
		go StartServer([]*RouteConfig{sc.PrivateRoutConfig}, privatePort, done)
		<-done
	}

	ExecuteInits(sc.Inits)

	<-done
}

func StartServer(routeConfig []*RouteConfig, port int, done chan bool) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	for i := range routeConfig {
		AddService(r, routeConfig[i])
	}

	r.Run(fmt.Sprintf(":%d", port))
	done <- true
}

func ReadConfigs(svcName string) {
	confPath := []string{
		fmt.Sprintf("%s/conf/", svcName),
		"conf/",
		"../conf/",
	}

	_ = MergeConfigIntoViper(confPath, "common.config", "yaml")

	log.Println("config read: ", viper.AllSettings())
}

func MergeConfigIntoViper(searchPaths []string, fileName, fileType string) error {
	f := fileName + "." + fileType
	merged := false
	for _, p := range searchPaths {
		viper.SetConfigFile(p + f)
		if viper.MergeInConfig() == nil {
			merged = true
			break
		}
	}

	if !merged {
		return errors.New("failed to read config file in any of the given search paths")
	}

	return nil
}

func AddService(r *gin.Engine, routeConfig *RouteConfig) {
	if routeConfig == nil {
		return
	}

	RegisterRoutes(r, *routeConfig)
}

func ExecuteInits(initFuncs []InitFunc) {
	for i := range initFuncs {
		initFuncs[i]()
	}
}
