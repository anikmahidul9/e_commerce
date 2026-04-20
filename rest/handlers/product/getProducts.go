package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"log"
	"net/http"

)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("GetProducts handler called")
	util.SendData(w,database.List(),http.StatusOK)
}	