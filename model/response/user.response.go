package response

import (
    "time"

)

type UserResponse struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"-" gorm:"column:password"`
	Phone string `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

}

type UserUpdateRequest struct {
	Name string `json:"name" validate:"required"`
	Phone string `json:"phone"` 
}


type UserUpdateEmailRequest struct {
	Email string `json:"email" validate:"required"`

}
