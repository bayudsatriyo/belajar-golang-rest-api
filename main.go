package main

import (
	"bayudsatriyo/belajar-golang-restfull-api/app"
	"bayudsatriyo/belajar-golang-restfull-api/controller"
	"bayudsatriyo/belajar-golang-restfull-api/helper"
	"bayudsatriyo/belajar-golang-restfull-api/middleware"
	"bayudsatriyo/belajar-golang-restfull-api/repository"
	router2 "bayudsatriyo/belajar-golang-restfull-api/router"
	"bayudsatriyo/belajar-golang-restfull-api/service"
	"fmt"
	validator2 "github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	db := app.NewDb()
	validator := validator2.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validator)
	categoryController := controller.NewCategoryController(categoryService)
	router := router2.NewCategoryRouter(categoryController)

	fmt.Println("Server is running on port 3000")
	fmt.Println("Running Golang REST API with MySQL Database 🚀")

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
