package services

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"tech-test-azura-lab/models"
	"tech-test-azura-lab/repository"
)

type SuccessResponse struct {
	Msg string `json:"msg"`
}

type ErrorResponse struct {
	Err string `json:"error"`
}

func GetAllProduct(productRepo repository.ProductRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		products, err := productRepo.GetAllProduct()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		body, err := json.Marshal(products)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})
}
func GetProductDetailByID(productRepo repository.ProductRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			resp, _ := json.Marshal(ErrorResponse{Err: "id is not found"})
			w.Write(resp)
			return
		}
		idNum, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp, _ := json.Marshal(ErrorResponse{Err: "id must be a number"})
			w.Write(resp)
			return
		}

		product, err := productRepo.GetProductDetail(idNum)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		body, err := json.Marshal(product)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})
}
func CreateProduct(productRepo repository.ProductRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp, _ := json.Marshal(ErrorResponse{Err: "Bad Request"})
			w.Write(resp)
			return
		}

		var product models.Product
		err = json.Unmarshal(body, &product)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = productRepo.CreateProduct(product)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp, _ := json.Marshal(SuccessResponse{Msg: "Success"})

		w.WriteHeader(http.StatusAccepted)
		w.Write(resp)
	})
}
func UpdateProduct(productRepo repository.ProductRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var newProduct models.Product
		err = json.Unmarshal(body, &newProduct)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		curProduct, err := productRepo.GetProductDetail(int(newProduct.ID))

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if newProduct.Nama != "" {
			curProduct.Nama = newProduct.Nama
		}
		if newProduct.Harga != 0 {
			curProduct.Harga = newProduct.Harga
		}
		if curProduct.Rating != 0 {
			curProduct.Rating = newProduct.Rating
		}
		if curProduct.Likes != 0 {
			curProduct.Likes = newProduct.Likes
		}
		if curProduct.Deskripsi != "" {
			curProduct.Deskripsi = newProduct.Deskripsi
		}

		err = productRepo.UpdateProduct(*curProduct)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp, _ := json.Marshal(SuccessResponse{Msg: "Success"})
		w.WriteHeader(http.StatusAccepted)
		w.Write(resp)
	})
}
func DeleteProduct(productRepo repository.ProductRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp, _ := json.Marshal(ErrorResponse{Err: "Bad Request"})
			w.Write(resp)
			return
		}

		var product models.Product
		err = json.Unmarshal(body, &product)
		if product.ID == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = productRepo.GetProductDetail(int(product.ID))
		if err != nil && err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = productRepo.DeleteProduct(int(product.ID))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp, _ := json.Marshal(SuccessResponse{Msg: "Delete Success"})
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
		return
	})
}
