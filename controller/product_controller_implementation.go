package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"tesMitraKasihPerkasa/helper"
	"tesMitraKasihPerkasa/model/web"
	"tesMitraKasihPerkasa/service"
)

type ProductControllerImplementation struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImplementation{
		ProductService: productService,
	}
}
func (controller ProductControllerImplementation) Insert(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {

	decoder := json.NewDecoder(request.Body)
	insertProduct := web.InputProductRequest{}
	err := decoder.Decode(&insertProduct)
	helper.PanicIfError(err)
	productResponse := controller.ProductService.Insert(request.Context(), insertProduct)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}
	helper.WriteToResponseBody(writter, webResponse)
}

func (controller ProductControllerImplementation) ListProduct(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productResponse := controller.ProductService.ListProduct(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}
	helper.WriteToResponseBody(writter, webResponse)
}

func (controller ProductControllerImplementation) EditStock(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	changeStock := web.EditStockRequest{}
	err := decoder.Decode(&changeStock)
	helper.PanicIfError(err)
	controller.ProductService.EditStock(request.Context(), changeStock)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writter, webResponse)
}
