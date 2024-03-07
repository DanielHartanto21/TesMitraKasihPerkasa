package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ProductController interface {
	Insert(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	ListProduct(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	EditStock(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
}
