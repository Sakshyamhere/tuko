package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"log/slog"
	"net/http"
	"sakshyahere/tuko/internal/config"
	"sakshyahere/tuko/internal/controller"
	"sakshyahere/tuko/internal/controller/auth"
	"sakshyahere/tuko/internal/di"
	"sakshyahere/tuko/internal/route"
)

func main() {
	// Echo instance
	config.LoadEnv()
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	registerRoutes(e)

	// Start server
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Register Routes
func registerRoutes(e *echo.Echo) {
	container := di.BuildContainer()
	err := di.InitializeApp(container)
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}
	// Inject controller and register routes
	err = container.Invoke(func(userController *controller.UserController, authController *auth.AuthController) {
		route.RegisterRoutes(e, userController, authController)
	})
	if err != nil {
		slog.Error("failed to register routes", "error", err)
	}
}
