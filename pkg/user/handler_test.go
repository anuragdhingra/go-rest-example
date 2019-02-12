package user

import (
	"bytes"
	"encoding/json"
	"github.com/anuragdhingra/go-rest-example/db/models"
	"github.com/go-chi/chi"
	"io"
	"net/http/httptest"
	"testing"
)

func TestSignup(t *testing.T) {
	router := chi.NewRouter()
	router.Post("/signup", Signup)

	user := createTestUser()

	request := httptest.NewRequest("POST","/signup", user)
	writer := httptest.NewRecorder()

	router.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestGetUsers(t *testing.T) {

	router := chi.NewRouter()
	router.Get("/users", GetUsers)

	request := httptest.NewRequest("GET","/users", nil)
	writer := httptest.NewRecorder()

	router.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	var users []models.User
	json.Unmarshal(writer.Body.Bytes(), &users)
	for _,user := range users {
		if user.Username != "testuser" {
			t.Error("Cannot retrieve JSON post")
		}
	}
}

func createTestUser() (jsonReader io.Reader) {
	user := &models.User{Username:"testuser", Email:"testemail@gmail.com", Password:"password"}
	userJSON, _ := json.Marshal(user)
	jsonReader = bytes.NewReader(userJSON)
	return jsonReader
}
