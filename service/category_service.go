package service

import (
	"bayudsatriyo/belajar-golang-restfull-api/model/web"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) (web.CategoryResponse, error)
	FindById(ctx context.Context, categoryId int) []web.CategoryResponse
}