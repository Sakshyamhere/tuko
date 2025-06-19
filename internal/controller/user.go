package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sakshyahere/tuko/internal/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (ctrl *UserController) GetUser(ctx echo.Context) error {
	message := ctrl.userService.GetUser()
	id := ctx.Get("user_id")
	res := map[string]interface{}{
		"id":      id,
		"message": message,
	}
	return ctx.JSON(http.StatusOK, res)
}
