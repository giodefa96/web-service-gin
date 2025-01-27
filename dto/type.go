package dto

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Body struct {
	Message string `json:"message"`
}
