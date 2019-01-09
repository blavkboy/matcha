package main

import (
	"net/http"
	"os"
	"time"

	"github.com/blavkboy/matcha/database"
	"github.com/blavkboy/matcha/mlogger"
	"github.com/blavkboy/matcha/routing"
	"github.com/blavkboy/matcha/services/auth"
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
	r.HandleFunc("/ws/{token}", routing.SocketConn)
	r.HandleFunc("/users/login", auth.NewToken(routing.HandleLogin)).Methods("POST")
	r.HandleFunc("/users/check", auth.ConfirmUser(routing.HandleCheck)).Methods("GET")
	r.HandleFunc("/users", routing.HandleUser).Methods("POST")
	r.HandleFunc("/users", auth.NewToken(routing.HandleUsers)).Methods("GET")
	r.HandleFunc("/home", routing.HandleHome).Methods("GET")
	r.HandleFunc("/", routing.HandleRoot).Methods("GET")
	http.ListenAndServe(":8080", r)
}
