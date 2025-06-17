package route

import (
	"github.com/labstack/echo/v4"
	"sakshyahere/tuko/internal/controller"
)

func RegisterRoutes(e *echo.Echo, userController *controller.UserController) {
	// You can group routes if needed
	v1 := e.Group("/api/v1")
	userGroup := v1.Group("/users")
	{
		userGroup.GET("/get", userController.GetUser)
		// Add more user-related routes here later
	}
}
