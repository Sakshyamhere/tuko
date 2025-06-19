package route

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"sakshyahere/tuko/internal/controller"
	"sakshyahere/tuko/internal/controller/auth"
	"sakshyahere/tuko/internal/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func RegisterRoutes(e *echo.Echo, userController *controller.UserController, authController *auth.AuthController) {
	// You can group routes if needed
	v1 := e.Group("/api/v1")
	e.Validator = &CustomValidator{validator: validator.New()}
	app := v1.Group("/app")
	app.Use(middleware.JWTMiddleware)

	//User Group
	userGroup := app.Group("/user")
	{
		userGroup.GET("/get", userController.GetUser)
	}

	//Auth Group
	authGroup := v1.Group("/auth")
	{
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/signup", authController.Signup)
		authGroup.GET("/email/xa", authController.EmailExists)
	}
}
