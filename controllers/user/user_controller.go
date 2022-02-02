package user

import (
	"immersive6/config"
	"immersive6/lib/database"
	"immersive6/models/base"
	_userDB "immersive6/models/user"
	_request "immersive6/models/user/request"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterController(c echo.Context) error {
	var userRegister _request.UserRegister
	c.Bind(&userRegister)

	// validasi request
	if userRegister.Email == "" {
		return c.JSON(http.StatusBadRequest, base.Response{
			Code:    http.StatusBadRequest,
			Message: "Email masih kosong",
			Data:    nil,
		})
	}

	if userRegister.Name == "" {
		return c.JSON(http.StatusBadRequest, base.Response{
			Code:    http.StatusBadRequest,
			Message: "Name masih kosong",
			Data:    nil,
		})
	}

	// validasi email

	userDB := userRegister.ToUserDB()
	var err error
	userDB, err = database.Register(userDB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.Response{
			Code:    http.StatusInternalServerError,
			Message: "Gagal register",
			Data:    nil,
		})
	}

	userResponse := userDB.ToUserResponse()
	return c.JSON(http.StatusOK, base.Response{
		Code:    http.StatusOK,
		Message: "Berhasil memasukkan ke DB",
		Data:    userResponse,
	})
}

func LoginController(c echo.Context) error {
	var userLogin _request.UserLogin
	c.Bind(&userLogin)

	var userDB _userDB.User
	result := config.DB.Where("email = ? AND password = ?", userLogin.Email, userLogin.Password).First(&userDB)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusUnauthorized, base.Response{
				Code:    http.StatusUnauthorized,
				Message: "Email dan password tidak sesuai",
				Data:    nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, base.Response{
			Code:    http.StatusInternalServerError,
			Message: "Gagal login",
			Data:    nil,
		})
	}

	responseUser := userDB.ToUserResponse()
	return c.JSON(http.StatusOK, base.Response{
		Code:    http.StatusOK,
		Message: "Success Login",
		Data:    responseUser,
	})
}
