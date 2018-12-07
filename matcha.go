package main

import (
	"net/http"

	"github.com/blavkboy/matcha/mlogger"
	"github.com/blavkboy/matcha/models"
	"github.com/blavkboy/matcha/routing"
	"github.com/gorilla/mux"
)

//In main we will handle all requests to the server
func main() {
	mlogger := mlogger.GetInstance()
	mlogger.Println("Starting 'Matcha' dating service API")
	models.Users = append(models.Users, *d)
	r := mux.NewRouter()
	//Register users to the users collectioon in the matcha database
	r.HandleFunc("/users", routing.HandleUser).Methods("POST")
	//Return all the users **Not yet implemented**
	r.HandleFunc("/users", routing.HandleUsers).Methods("GET")
	http.ListenAndServe(":8080", r)
}
