package user

import (
	"encoding/json"
	"fmt"
	"github.com/anuragdhingra/go-rest-example/db"
	"github.com/anuragdhingra/go-rest-example/db/models"
	"github.com/anuragdhingra/go-rest-example/pkg/helper"
	"net/http"
)

type Request struct {
	Username string
	Email string
	Password string
}

type Response struct {
	Id uint `json:"id"`
	Message string `json:"message"`
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

	response := Response{
		Id:user.ID,
		Message:"User created successfully!",
	}
	helper.RespondwithJSON(w, http.StatusOK,response)
	return
}