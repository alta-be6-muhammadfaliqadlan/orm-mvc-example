package request

import "immersive6/models/user"

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Country  string `json:"country"`
}

func (u *UserRegister) ToUserDB() user.User {
	return user.User{
		Name: u.Name,
		Email: u.Email,
		Password: u.Password,
		Address: u.Address,
		Country: u.Country,
	}
}