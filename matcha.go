package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
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

//this data struct will hold all the dependencies
//needed to provide our web service
type data struct {
	PageTitle string
}

//In main we will handle all requests to the server
func main() {
	http.HandleFunc("/", handleRoot)
	http.ListenAndServe(":8080", nil)
}
