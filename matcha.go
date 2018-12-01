package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/blavkboy/matcha/models"
	"github.com/gorilla/mux"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	sampleData := data{
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

//handleUsers will only be used when a parameter is passed to the handler function
func handleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	for _, u := range users {
		if u.ID == vars["id"] {
			json.NewEncoder(w).Encode(u)
		}
	}
}

//handle User will handle requests to get the users from the browser.
func handleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//we process get request and return either the selected user or
	//all the users
	//Todo: refine search capabilities and make this more efficient
	if r.Method == "GET" {
		if vars["id"] != "" {
			handleUsers(w, r)
			return
		}
		tmpl, err := template.ParseFiles("views/root.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		tmpl.Execute(w, data{
			PageTitle: "Test Page",
			Data:      users,
		})
	} else if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var body models.User
		json.NewDecoder(r.Body).Decode(&body)
		users = append(users, body)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

//this data struct will hold all the dependencies
//needed to provide our web service
type data struct {
	PageTitle string
	Data      []models.User
}

//In main we will handle all requests to the server
func main() {
	users = append(users, *d)
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/static/", fs))
	r.HandleFunc("/", handleRoot)
	r.HandleFunc("/users", handleUser).Methods("GET", "POST")
	r.HandleFunc("/users/{id}", handleUser).Methods("GET")
	http.ListenAndServe(":8080", r)
}
