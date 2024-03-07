package web

type UserLoginRequest struct {
	Name     string `validate:"required"  json:"name"`
	Password string `validate:"required"  json:"password"`
}
