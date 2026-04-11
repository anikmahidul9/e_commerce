package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"log"
	"net/http"

)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("GetProducts handler called")
	util.SendData(w,database.List(),http.StatusOK)
}	