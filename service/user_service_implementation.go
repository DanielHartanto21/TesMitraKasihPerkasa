package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"tesMitraKasihPerkasa/helper"
	"tesMitraKasihPerkasa/model/domain"
	"tesMitraKasihPerkasa/model/web"
	"tesMitraKasihPerkasa/repository"
)

type UserServiceImplementation struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImplementation{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImplementation) Logins(ctx context.Context, request web.UserLoginRequest) web.LoginResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	user := domain.Users{
		NamaUser: request.Name,
		Password: request.Password,
	}
	user, err = service.UserRepository.Logins(ctx, tx, user)
	helper.PanicIfError(err)
	return helper.ToLogin(user)
}
func (service *UserServiceImplementation) Check(uuid string) bool {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	return service.UserRepository.Check(tx, uuid)
}
