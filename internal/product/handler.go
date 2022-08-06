package product

import (
	"encoding/json"
	"muxtemp/internal/db"
	"muxtemp/internal/entity"
	"net/http"
	"github.com/gorilla/mux"
	"muxtemp/pkg/log"
	"go.uber.org/zap"
	"time"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product entity.Product
	json.NewDecoder(r.Body).Decode(&product)
	db.Instance.Create(&product)
	json.NewEncoder(w).Encode(product)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("GetProductById",
		zap.Duration("duration", time.Millisecond),
		zap.Int("size", int(r.ContentLength)),
		zap.String("content-type", "application/json"),
		zap.String("user-agent", r.UserAgent()),
		zap.Int("status", http.StatusOK),
		zap.String("method", r.Method),
		zap.String("uri-path", r.URL.Path),
		zap.String("host", r.Host),
		zap.String("params", mux.Vars(r)["id"]),
		zap.String("client", r.RemoteAddr),
	)
	productId := mux.Vars(r)["id"]
	if checkIfProductExists(productId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entity.Product
	db.Instance.First(&product, productId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("GetProducts",
		zap.Int("size", int(r.ContentLength)),
		zap.String("content-type", "application/json"),
		zap.String("user-agent", r.UserAgent()),
		zap.Int("status", http.StatusOK),
		zap.String("method", r.Method),
		zap.String("params", r.URL.Path),
		zap.String("client", r.RemoteAddr),
	)
	var products []entity.Product
	db.Instance.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if checkIfProductExists(productId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entity.Product
	db.Instance.First(&product, productId)
	json.NewDecoder(r.Body).Decode(&product)
	db.Instance.Save(&product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	productId := mux.Vars(r)["id"]
	if checkIfProductExists(productId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entity.Product
	db.Instance.Delete(&product, productId)
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
}

func checkIfProductExists(productId string) bool {
	var product entity.Product
	db.Instance.First(&product, productId)
	if product.ID == 0 {
		return false
	}
	return true
}
