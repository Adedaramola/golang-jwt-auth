package main

import (
	"net/http"

	"github.com/adedaramola/golang-jwt-auth/auth"
	"github.com/adedaramola/golang-jwt-auth/utils"
	"github.com/julienschmidt/httprouter"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request LoginRequest

	err := utils.ShouldBindJSON(r, &request)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, utils.H{"error": "bad request"})
	}
	// validate request fields

	// check user against database records

	// verify password hash

	// everything is fine, generate token
	token, err := auth.GenerateToken(&auth.Payload{
		Email: request.Email,
	})
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, utils.H{"error": "Server error"})
		return
	}

	utils.JSON(w, http.StatusOK, utils.H{"token": token})
	return
}

func Protected(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	utils.JSON(w, http.StatusOK, utils.H{"message": "protected resource"})
	return
}

func PingServer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	utils.JSON(w, http.StatusOK, utils.H{"message": "server is alive"})
	return
}
