package server

import (
	"github.com/labstack/echo/v4"
	"github.com/tzilist/m-service/config"
)

func beforeCreateServer() {
	config.InitConfigs()
}

func afterCreateServer() {
}

// Initializes a new server
func Start() *echo.Echo {
	beforeCreateServer()

	server := echo.New()

	return server
}
