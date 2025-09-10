package dto

type UserCreateDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
