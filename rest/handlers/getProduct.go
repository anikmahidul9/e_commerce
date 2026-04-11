package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)


func GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	pId,err := strconv.Atoi(productID)
	if err !=nil {
		util.SendError(w, "Invalid product ID", http.StatusBadRequest)
	}
	product := database.GetById(pId)
	if product == nil {
		util.SendError(w, "Product not found", http.StatusNotFound)
		return
	}
	util.SendData(w,product,http.StatusOK)
}