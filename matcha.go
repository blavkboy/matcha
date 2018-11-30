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

//array we will keep users in
var users []models.User

//sample user to add to the array
var d = &models.User{
	ID:       "1",
	Username: "The dark one",
	Fname:    "El Pharoah",
	Lname:    "",
	Email:    "akjhdskjahsdkjh@sponges.com",
	Password: "akjsdghhjakhdksjadklashdklad",
	Bio: &models.Bio{
		Caddress: nil,
		Oaddress: &models.Address{
			Street1:  "352",
			Street2:  "Du Toit Street",
			Suburb:   "Wierda Park",
			City:     "Pretoria",
			Province: "Gauteng",
		},
		Sexuality: &models.Sexuality{
			Sex:         models.Male,
			Orientation: models.Hetero,
			Looking:     models.Fun,
			Preferences: nil,
		},
		Hobbies:   nil,
		Interests: nil,
	},
}

//handle Users will handle requests to get the users from the browser.
func handleUser(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	if r.Method == "GET" {
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
