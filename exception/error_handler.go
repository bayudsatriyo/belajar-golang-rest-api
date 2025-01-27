package exception

import (
	"bayudsatriyo/belajar-golang-restfull-api/helper"
	"bayudsatriyo/belajar-golang-restfull-api/model/web"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	execption, ok := err.(validator.ValidationErrors)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponses := web.WebResponses{
			Code:   http.StatusBadRequest,
			Status: "Bad Request Error",
			Data:   execption.Error(),
		}
		helper.WriteToResponseBody(writer, webResponses)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	execption, ok := err.(NotFoundError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponses := web.WebResponses{
			Code:   http.StatusNotFound,
			Status: "Not Found Error",
			Data:   execption.Error,
		}
		helper.WriteToResponseBody(writer, webResponses)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponses := web.WebResponses{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}
	helper.WriteToResponseBody(writer, webResponses)
}
