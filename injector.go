//go:build wireinject
// +build wireinject

package main

import (
	"bayudsatriyo/belajar-golang-restfull-api/app"
	"bayudsatriyo/belajar-golang-restfull-api/controller"
	"bayudsatriyo/belajar-golang-restfull-api/middleware"
	"bayudsatriyo/belajar-golang-restfull-api/repository"
	"bayudsatriyo/belajar-golang-restfull-api/router"
	"bayudsatriyo/belajar-golang-restfull-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func NewValidator() *validator.Validate {
	return validator.New()
}

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDb,
		NewValidator,
		categorySet,
		router.NewCategoryRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
