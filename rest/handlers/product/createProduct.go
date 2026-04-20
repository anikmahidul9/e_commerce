package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"log"
	"net/http"
	"fmt"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdProduct := database.Store(newProduct)
	util.SendData(w, createdProduct, http.StatusCreated)
	log.Printf("Product created: %s", fmt.Sprintf("%+v", createdProduct))
}
