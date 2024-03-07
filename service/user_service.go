package service

import (
	"context"
	"tesMitraKasihPerkasa/model/web"
)

type UserService interface {
	Login(ctx context.Context, request web.UserLoginRequest) string
}
