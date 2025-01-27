package helper

import (
	"bayudsatriyo/belajar-golang-restfull-api/model/domain"
	"bayudsatriyo/belajar-golang-restfull-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
