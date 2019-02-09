package user

import (
	"encoding/json"
	"fmt"
	"github.com/anuragdhingra/go-rest-example/db"
	"github.com/anuragdhingra/go-rest-example/db/models"
	"github.com/anuragdhingra/go-rest-example/pkg/helper"
	"net/http"
)

type signupResponse struct {
	Id uint `json:"id"`
	Message string `json:"message"`
}

type publicUser struct {
	Username string `json:"username"`
	Email string `json:"email"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	if dbc:= db.GetDb().Create(&user); dbc.Error != nil {
		helper.RespondWithError(w, http.StatusBadRequest, dbc.Error)
		return
	}

	response := signupResponse{
		Id:user.ID,
		Message:"User created successfully!",
	}
	helper.RespondwithJSON(w, http.StatusOK,response)
	return
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []models.User{}
	if dbc:= db.GetDb().Find(&users); dbc.Error != nil {
		helper.RespondWithError(w, http.StatusBadRequest, dbc.Error)
		return
	}
	publicUsers := []publicUser{}
	for _,user := range users {
		publicUser := publicUser{Username:user.Username, Email:user.Email}
		publicUsers = append(publicUsers, publicUser)
	}
	helper.RespondwithJSON(w, http.StatusOK,publicUsers)
	return
}