package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		util.SendError(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
     database.Delete(pId)
	
	util.SendData(w, map[string]string{"message": "Product deleted successfully"}, http.StatusOK)
}