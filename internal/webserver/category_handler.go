package webserver

import "github.com/Viniciuscpires/fc_api_product/internal/service"

type WebCategoryServer struct {
	CategoryService *service.CategoryService
}

func NewWebCategoryServer(categoryService *service.CategoryService) *WebCategoryServer {
	return &WebCategoryServer{
		CategoryService: categoryService,
	}
}
