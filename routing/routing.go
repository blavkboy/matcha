package routing

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/blavkboy/matcha/mlogger"
	"github.com/blavkboy/matcha/models"
	"github.com/gorilla/mux"
)

type middleWare func(next http.HandlerFunc) http.HandlerFunc

//HandleRoot will handle calls to the root of the domain
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sampleData := models.Data{}
	if strings.Compare(r.Method, "GET") == 0 {
		json.NewEncoder(w).Encode(sampleData)
	}
}

//HandleUsers will only be used when a parameter is passed to the handler function
func HandleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	for _, u := range models.Users {
		if u.ID == vars["id"] {
			json.NewEncoder(w).Encode(u)
		}
	}
}

//HandleUser will handle requests to get the users from the browser.
func HandleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mlogger := mlogger.GetInstance()
	vars := mux.Vars(r)
	//we process get request and return either the selected user or
	//all the users
	//Todo: refine search capabilities and make this more efficient
	if r.Method == "GET" {
		mlogger.Println("Recieved 'GET' request from: " + r.UserAgent())
		if vars["id"] != "" {
			HandleUsers(w, r)
			return
		}
		json.NewEncoder(w).Encode(models.Users)
	} else if r.Method == "POST" {
		var body models.User
		json.NewDecoder(r.Body).Decode(&body)
		models.Users = append(models.Users, body)
	}
}
