package user

import (
	"bytes"
	"encoding/json"
	"github.com/anuragdhingra/go-rest-example/db/models"
	"github.com/go-chi/chi"
	"net/http/httptest"
	"testing"
)

func TestSignup(t *testing.T) {
	router := chi.NewRouter()
	router.Post("/signup", Signup)

	user := &models.User{Username:"testuser", Email:"testemail@gmail.com", Password:"password"}
	userJSON, _ := json.Marshal(user)
	jsonReader := bytes.NewReader(userJSON)

	request := httptest.NewRequest("POST","/signup", jsonReader)
	writer := httptest.NewRecorder()

	router.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
