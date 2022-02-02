package database

import (
	"immersive6/config"
	"immersive6/models/user"
)

func Register(userDB user.User) (user.User, error) {
	result := config.DB.Create(&userDB)

	if result.Error != nil {
		return user.User{}, result.Error
	}

	return userDB, nil
}
