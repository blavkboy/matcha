package routing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/blavkboy/matcha/mlogger"
	"github.com/blavkboy/matcha/models"
	"github.com/blavkboy/matcha/services/auth"
	"github.com/blavkboy/matcha/views"
	"golang.org/x/crypto/bcrypt"
)

type middleWare func(next http.HandlerFunc) http.HandlerFunc

//HandleRoot will handle calls to the root of the domain
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	views.RenderIndex(w)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	//something todo later
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
	pass, err := bcrypt.GenerateFromPassword([]byte(body.Password), models.Cost)
	if err != nil {
		fmt.Println("Error hashing password: ", err.Error())
		return
	}
	body.Password = string(pass)
	mlogger.Println("Saving user with password: ", string(body.Password))
	newBody := models.FindUser("email", body.Email)
	if newBody != nil {
		fmt.Fprintf(w, "Fail email")
		return
	}
	newBody = &models.User{}
	newBody = models.FindUser("username", body.Username)
	if newBody != nil {
		fmt.Fprintf(w, "Fail username")
		return
	}
	ret := models.NewUser(&body)
	if ret == nil {
		fmt.Fprintf(w, "Fail")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(ret)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		mlogger.Println("Error: ", err)
		return
	}
	mlogger.Println(time.Now())
	mlogger.Println(body)
}

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func HandleCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Something")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	views.RenderHome(w, auth.GetCurrentUser(r))
}
