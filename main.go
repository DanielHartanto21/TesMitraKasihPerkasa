package main

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"tesMitraKasihPerkasa/app"
	"tesMitraKasihPerkasa/controller"
	"tesMitraKasihPerkasa/helper"
	"tesMitraKasihPerkasa/middleware"
	"tesMitraKasihPerkasa/repository"
	"tesMitraKasihPerkasa/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)
	router := app.NewRouter(userController, productController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
