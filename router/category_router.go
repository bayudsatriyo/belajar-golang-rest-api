package router

import (
	"bayudsatriyo/belajar-golang-restfull-api/controller"
	"bayudsatriyo/belajar-golang-restfull-api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewCategoryRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:id", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PATCH("/api/categories/:id", categoryController.Update)
	router.DELETE("/api/categories/:id", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
