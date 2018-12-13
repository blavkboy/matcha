package routing

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/blavkboy/matcha/mlogger"
	"github.com/blavkboy/matcha/models"
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

//HandleUser will handle requests to get the users from the browser.
//we can abstract some of it to make the login method and let the user
//keep his/her state using the token
func HandleUser(w http.ResponseWriter, r *http.Request) {
	mlogger := mlogger.GetInstance()
	//we process get request and return either the selected user or
	//all the users
	//Todo: refine search capabilities and make this more efficient
	var body models.User
	json.NewDecoder(r.Body).Decode(&body)
	mlogger.Println("Saving user")
	models.NewUser(&body)
	mlogger.Println(time.Now())
	mlogger.Println(body)
}

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
