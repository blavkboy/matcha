package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gmohlamo/matcha/mlogger"
	"github.com/gmohlamo/matcha/models"
	"golang.org/x/crypto/bcrypt"
)

var mySigningKey = []byte("The `jig is up")

//Login function things
func Login(w http.ResponseWriter, r *http.Request) {
	logger := mlogger.GetInstance()
	fmt.Printf("%v\n", r.Body)
	fmt.Println(r.Method + "\n")
	session := GetSession(r)
	if session != nil {
		session.Options.MaxAge = -1
	}
	if strings.Compare(r.Method, "POST") == 0 {
		var user models.User
		var compare *models.User
		json.NewDecoder(r.Body).Decode(&user)
		fmt.Printf("Decoded request body --> %v\n", string(user.Username) != "")
		if string(user.Username) == "" || string(user.Password) == "" {
			w.Write([]byte("{\"success\": false}"))
			return
		}
		// Create a new token object, specifying signing method and the claims
		// you would like it to contain.
		compare = models.FindUser("username", user.Username)
		if compare == nil {
			w.Write([]byte("{\"success\": false}"))
			return
		}
		pass1 := []byte(compare.Password)
		pass2 := []byte(user.Password)
		if bcrypt.CompareHashAndPassword(pass1, pass2) != nil {
			w.Write([]byte("{\"success\": false}"))
			return
		}
		logger.Println("Password match for user: ", bcrypt.CompareHashAndPassword(pass1, pass2) == nil)
		logger.Println("DEBUG INFO user_id: ", compare.ID)
		// write the user body to the stream
		session := GetSession(r)
		session.Options.MaxAge = 3600 * 24 * 3
		fmt.Println(session)
		if session == nil {
			w.Write([]byte("{\"success\": false}"))
			return
		}
		session.Values["user"] = string(compare.ID)
		compare.Password = ""
		fmt.Println(session)
		err := session.Save(r, w)
		if err != nil {
			logger.Println(err)
			w.Write([]byte("{\"success\": false}"))
			return
		}
		body, err := json.Marshal(compare)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"success\": false}"))
			return
		}
		w.Write([]byte("{\"success\": true," + "\"user\":" + string(body) + "}"))
	}
}

//ConfirmUser Should inform the client of the user who is currently logged in
func ConfirmUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := GetSession(r)
		if session == nil {
			w.Write([]byte("{\"success\": false}"))
		} else {
			next(w, r)
		}
	}
}

//GetCurrentUser basically a redundant version of confirm user
func GetCurrentUser(r *http.Request) *models.User {
	session := GetSession(r)
	username := session.Values["user"]
	if username == nil {
		return nil
	}
	return models.FindUser("username", username)
}

func failedAuth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{\"success\": false}"))
}
