package request

type UserCreateRequest struct {
	Name string `json:"name" validate:"required"` 
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Phone string `json:"phone"`
}

type UserUpdateRequest struct {
	Name string `json:"name" validate:"required"`
	Phone string `json:"phone"` 
}

type UserUpdateEmailRequest struct {
	Email string `json:"email" validate:"required"`

}

type UserUpdatePassword struct {

}