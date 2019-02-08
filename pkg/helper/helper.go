package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Catch(err error) {
	if err != nil {
		panic(err)
	}
}

func RespondWithError(w http.ResponseWriter, code int, err error) {
	RespondwithJSON(w, code, err)
}

func RespondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}