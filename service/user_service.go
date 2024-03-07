package service

import (
	"context"
	"tesMitraKasihPerkasa/model/web"
)

type UserService interface {
	Logins(ctx context.Context, request web.UserLoginRequest) web.LoginResponse
	Check(uuid string) bool
}
