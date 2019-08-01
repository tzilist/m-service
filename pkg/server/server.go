package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tzilist/m-service/config"
	"github.com/tzilist/m-service/pkg/router"
	"github.com/tzilist/m-service/pkg/util/validator"
)

type (
	// Hook hook for before and after creating the server struct
	Hook func()
)

// beforeCreateServer inits anything that needs to happen before the server can start
func beforeCreateServer() {
	config.InitConfigs()
}

// Start a new server
// Accepts a before/after create hook (which may be useful for testing)
func Start(beforeCreateHook Hook, afterCreateHook Hook) (*echo.Echo, string) {
	// allow custom before create server hook
	if beforeCreateHook != nil {
		beforeCreateHook()
	}

	beforeCreateServer()

	server := echo.New()
	// hide some nonsense that echo uses on startup
	server.HideBanner = true
	server.HidePort = true

	// set middlewares
	server.Use(middleware.Recover())
	server.Use(middleware.Gzip())
	server.Use(middleware.Logger())

	server.Validator = validator.CreateValidator()

	// set routes
	router.CreateRoutes(server)

	serverPort := fmt.Sprintf(":%d", config.Config.Server.Port)

	// after create server hooks
	if afterCreateHook != nil {
		afterCreateHook()
	}

	return server, serverPort
}
