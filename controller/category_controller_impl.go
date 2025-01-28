package controller

import (
	"bayudsatriyo/belajar-golang-restfull-api/helper"
	"bayudsatriyo/belajar-golang-restfull-api/model/web"
	"bayudsatriyo/belajar-golang-restfull-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponses{
		Code:   http.StatusCreated,
		Status: "Success",
		Data:   categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponses{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.Delete(request.Context(), id)
	webResponse := web.WebResponses{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse, err := controller.CategoryService.FindById(request.Context(), id)

	helper.PanicIfError(err)

	webResponse := web.WebResponses{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   categoryResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponses{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   categoryResponses,
	}

	writer.Header().Add("Content-Type", "application/json")
	helper.WriteToResponseBody(writer, webResponse)
}
