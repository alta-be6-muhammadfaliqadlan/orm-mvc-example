package user

import (
	"immersive6/models/user/response"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Name      string         `json:"name" gorm:"not null"`
	Email     string         `json:"email" form:"email" gorm:"unique"`
	Password  string         `json:"password" form:"password"`
	Address   string         `json:"address" form:"address"`
	Country   string         `json:"country" form:"country" gorm:"type:VARCHAR(10)"`
}

func (u *User) ToUserResponse() response.UserResponse{
	return response.UserResponse{
		ID: u.ID,
		Email: u.Email,
		Name: u.Name,
		Address: u.Address,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
