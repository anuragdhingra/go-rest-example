package main

import (
	"fmt"
	_ "github.com/anuragdhingra/go-rest-example/db"
	"github.com/anuragdhingra/go-rest-example/pkg/user"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"time"
)
var router *chi.Mux
func main() {
	router = chi.NewRouter()

	router.Post("/signup", user.Signup)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", Logger()))
}

func Logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.Method, r.URL)
		router.ServeHTTP(w,r)
	})
}