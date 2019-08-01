package router

import (
	"github.com/labstack/echo/v4"
	"github.com/tzilist/m-service/pkg/controllers"
)

// CreateRoutes creates the server routes
func CreateRoutes(server *echo.Echo) {
	server.GET("/:channel/messages", controllers.GetChannelMessages)
	server.POST("/:channel/messages", controllers.PostChannelMessage)
}
