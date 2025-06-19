package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sakshyahere/tuko/internal/service/auth"
)

type (
	AuthController struct {
		service auth.AuthService
	}
	User struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required"`
		Token     string `json:"token"`
	}
	Email struct {
		Email string `json:"email" validate:"required,email"`
	}
)

func NewAuthController(service auth.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (ctrl *AuthController) Login(ctx echo.Context) error {
	user := new(User)
	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	if err := ctx.Validate(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	token, err := ctrl.service.LoginService(user.Email, user.Password)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	res := map[string]interface{}{
		"email": user.Email,
		"token": token,
	}
	return ctx.JSON(http.StatusOK, res)
}

func (ctrl *AuthController) Signup(ctx echo.Context) error {
	user := new(User)
	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	if err := ctx.Validate(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	token, err := ctrl.service.SignupService(user.Email, user.FirstName, user.LastName, user.Password)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	res := map[string]interface{}{
		"email": user.Email,
		"token": token,
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (ctrl *AuthController) EmailExists(ctx echo.Context) error {
	email := new(Email)
	if err := ctx.Bind(email); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	if err := ctx.Validate(email); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	err := ctrl.service.EmailExistsService(email.Email)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	res := map[string]interface{}{
		"message": "Email Available",
	}
	return ctx.JSON(http.StatusCreated, res)
}
