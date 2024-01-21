package main

import (
	"github.com/pedrosandrini/gopportunities/config"
	router "github.com/pedrosandrini/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// initialize config
	err := config.Init()
	if err != nil {
		logger.ErrorF("config initialization error: %v", err)
		return
	}

	// Initializae routers
	router.Initialize()

}
