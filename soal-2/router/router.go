package router

import (
	"net/http"
	"tech-test-azura-lab/middleware"
	"tech-test-azura-lab/repository"
	"tech-test-azura-lab/services"
)

func GenerateMux(productRepo *repository.ProductRepo) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/api/product/all",
		middleware.Method("GET", middleware.ContentType("application/json", services.GetAllProduct(*productRepo))))
	mux.Handle("/api/product",
		middleware.Method("GET", middleware.ContentType("application/json", services.GetProductDetailByID(*productRepo))))
	mux.Handle("/api/product/update",
		middleware.Method("PUT", middleware.ContentType("application/json", services.UpdateProduct(*productRepo))))
	mux.Handle("/api/product/create",
		middleware.Method("POST", middleware.ContentType("application/json", services.CreateProduct(*productRepo))))
	mux.Handle("/api/product/delete",
		middleware.Method("DELETE", middleware.ContentType("application/json", services.DeleteProduct(*productRepo))))

	return mux
}
