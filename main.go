package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"rochiekop/golang-restful-api/app"
	"rochiekop/golang-restful-api/controller"
	"rochiekop/golang-restful-api/helper"
	"rochiekop/golang-restful-api/middleware"
	"rochiekop/golang-restful-api/repository"
	"rochiekop/golang-restful-api/service"
)

func main() {
	
	validate := validator.New()
	db := app.NewDB()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
