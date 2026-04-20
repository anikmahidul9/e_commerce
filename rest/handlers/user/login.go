package user

import (
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	user,err := database.Find(loginReq.Email, loginReq.Password)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	util.SendData(w, user, http.StatusOK)
	log.Printf("User logged in: %s", fmt.Sprintf("%+v", user))

	cnf := config.GetConfig()
	
	accessToken, err := util.CreateJwt(cnf.SecretKey, util.JwtPayload{
		Sub:         fmt.Sprintf("%d", user.ID),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopOwner: user.IsShopOwner,
	})
	if err != nil {
		http.Error(w, "Failed to create access token", http.StatusInternalServerError)
		return
	}

	util.SendData(w, map[string]string{"access_token": accessToken}, http.StatusOK)


}