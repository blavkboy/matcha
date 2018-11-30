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

//handle Users will handle requests to get the users from the browser.
func handleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	d := &models.User{ID: "skjhgdjahds", Username: "The dark one", Fname: "El Pharoah", Email: "akjhdskjahsdkjh@sponges.com", Password: "akjsdghhjakhdksjadklashdklad"}
	err := json.NewEncoder(w).Encode(d)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//this data struct will hold all the dependencies
//needed to provide our web service
type data struct {
	PageTitle string
}

//In main we will handle all requests to the server
func main() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/static/", fs))
	r.HandleFunc("/", handleRoot)
	r.HandleFunc("/users/{id}", handleUser)
	http.ListenAndServe(":8080", r)
}
