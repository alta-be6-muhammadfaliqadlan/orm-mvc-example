package route

import (
	"immersive6/controllers/user"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.POST("users/register", user.RegisterController)
	e.POST("users/login", user.LoginController)
	return e
}
