package app

import (
	"github.com/julienschmidt/httprouter"
	"tesMitraKasihPerkasa/controller"
	"tesMitraKasihPerkasa/exception"
)

func NewRouter(userController controller.UserController, productRouter controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/login", userController.Logins)
	router.POST("/insert_product", productRouter.Insert)
	router.GET("/list_product", productRouter.ListProduct)
	router.PATCH("/update_stock", productRouter.EditStock)
	router.PanicHandler = exception.ErrorHandler

	return router
}
