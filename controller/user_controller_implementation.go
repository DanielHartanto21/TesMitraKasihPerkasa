package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"tesMitraKasihPerkasa/helper"
	"tesMitraKasihPerkasa/model/web"
	"tesMitraKasihPerkasa/service"
)

type UserControllerImplementation struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImplementation{
		UserService: userService,
	}
}
func (controller *UserControllerImplementation) Logins(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	decoder := json.NewDecoder(request.Body)
	userLogin := web.UserLoginRequest{}
	err := decoder.Decode(&userLogin)
	helper.PanicIfError(err)
	userResponse := controller.UserService.Logins(request.Context(), userLogin)
	uuid := userResponse.Uuid
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data: struct {
			AccessToken string `json:"uuid"`
		}{
			AccessToken: uuid,
		},
	}
	writter.Header().Add("Authorization", uuid)
	helper.WriteToResponseBody(writter, webResponse)
}
