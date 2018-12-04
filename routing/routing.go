package routing

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/blavkboy/matcha/mlogger"
	"github.com/blavkboy/matcha/models"
	"github.com/gorilla/mux"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	sampleData := models.Data{
		PageTitle: "Root page",
	}
	if strings.Compare(r.Method, "GET") == 0 {
		tmpl, err := template.ParseFiles("views/root.html")
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		tmpl.Execute(w, sampleData)
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
		tmpl, err := template.ParseFiles("views/root.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		tmpl.Execute(w, models.Data{
			PageTitle: "Test Page",
			Data:      models.Users,
		})
	} else if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var body models.User
		json.NewDecoder(r.Body).Decode(&body)
		models.Users = append(models.Users, body)
	}
}
