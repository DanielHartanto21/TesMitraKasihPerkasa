package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController interface {
	Logins(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
}
