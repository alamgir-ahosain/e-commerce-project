package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alamgir-ahosain/e-commerce-project/internal/models"
)

var userList = []models.User{}

// make JSON format
func MakeJSONFormatFuncUser(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(userList)
}

// create product function:post request
func CreateUserFunc(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Give the valid json format", 400)
		return
	}

	if newUser.ID != 0 {
		http.Error(w, "Duplicate User", 400)
		return
	}
	newUser.ID = len(userList) + 1
	userList = append(userList, newUser)
	MakeJSONFormatFuncUser(w, 201)
}

// Login User
func LoginUserFunc(w http.ResponseWriter, r *http.Request) (models.User, error) {

	var reqLogin models.RequestLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		return models.User{}, fmt.Errorf("Give the valid json format")
	}

	for i, val := range userList {
		if val.Email == reqLogin.Email && val.Password == reqLogin.Password {
			return userList[i], nil
		}
	}
	return models.User{}, fmt.Errorf("user not found")

}
