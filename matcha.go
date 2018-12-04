package main

import (
	"net/http"

	"github.com/blavkboy/matcha/mlogger"
	"github.com/blavkboy/matcha/models"
	"github.com/blavkboy/matcha/routing"
	"github.com/blavkboy/matcha/services/auth"
	"github.com/gorilla/mux"
)

//In main we will handle all requests to the server
func main() {
	mlogger := mlogger.GetInstance()
	mlogger.Println("Starting 'Matcha' dating service API")
	models.Users = append(models.Users, *d)
	r := mux.NewRouter()
	r.HandleFunc("/", auth.NewToken(routing.HandleRoot)).Methods("GET")
	r.HandleFunc("/users", auth.NewToken(routing.HandleUser)).Methods("GET", "POST")
	r.HandleFunc("/users/{id}", auth.NewToken(routing.HandleUser)).Methods("GET")
	http.ListenAndServe(":8080", r)
}
