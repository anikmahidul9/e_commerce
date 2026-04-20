package user

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	createdUser, err := newUser.Store()
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	util.SendData(w, createdUser, http.StatusCreated)
	log.Printf("User created: %s", fmt.Sprintf("%+v", createdUser))
}
