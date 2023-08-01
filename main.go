package main

import (
	"github.com/ahmetson/sample-extension/handler"
	"github.com/ahmetson/service-lib/configuration"
	"github.com/ahmetson/service-lib/extension"
	"github.com/ahmetson/service-lib/log"
)

func main() {
	logger, _ := log.New("sample", false)
	appConfig, _ := configuration.New(logger)

	service, _ := extension.New(appConfig, logger)
	service.AddController(configuration.ReplierType)

	service.RequireProxy("github.com/ahmetson/web-proxy", configuration.DefaultContext)
	service.Pipe("github.com/ahmetson/web-proxy", service.GetControllerName())

	service.GetController().AddRoute(handler.Add())
	service.GetController().AddRoute(handler.Mul())

	err := service.Prepare()
	if err != nil {
		logger.Fatal("service.Prepare", "error", err)
	}

	service.Run()
}
