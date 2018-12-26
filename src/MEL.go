package main

import (
	"api"
	"config"
	"pandora"
	"pandora/modules/logger"
	"services"
)

func main() {
	logger.D("starting")
	pandora.Init("cfg.conf")
	config.Init()
	services.Init()
	pandora.RouteApi("/jsonApi", api.NewCommonApi())
	pandora.Start()
}
