package main

import (
	"net/http"
	"os"
	"time"

	"github.com/blavkboy/matcha/database"
	"github.com/blavkboy/matcha/mlogger"
	"github.com/blavkboy/matcha/routing"
	"github.com/blavkboy/matcha/services/auth"
	"github.com/blavkboy/matcha/views"
	"github.com/gorilla/mux"
)

//In main we will handle all requests to the server
func main() {
	mlogger := mlogger.GetInstance()
	mlogger.Println(time.Now())
	err, conn := database.InitDB()
	defer conn.Close()
	if err != nil {
		mlogger.Println("Error: ", err)
		os.Exit(1)
	}
	mlogger.Println("Starting 'Matcha' dating service API")
	r := mux.NewRouter()

	//Register users to the users collectioon in the matcha database
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.HandleFunc("/users", routing.HandleUser).Methods("POST")
	r.HandleFunc("/users", auth.NewToken(routing.HandleUsers)).Methods("GET")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		views.RenderIndex(w)
	})
	http.ListenAndServe(":8080", r)
}
