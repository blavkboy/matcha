package routing

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/blavkboy/matcha/mlogger"
	"github.com/blavkboy/matcha/models"
	"github.com/blavkboy/matcha/services/auth"
	"github.com/blavkboy/matcha/socket"
	"github.com/blavkboy/matcha/views"
	"github.com/gorilla/websocket"
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

func SocketConn(w http.ResponseWriter, r *http.Request) {
	var start int = 0
	mlogger := mlogger.GetInstance()
	token := strings.Split(r.URL.Path, "ws/")[1]
	if token == "" {
		mlogger.Println("Error getting jwt string")
		fmt.Fprint(w, "Connection rejected")
		return
	}
	user := auth.GetUserFromString(token)
	if user == nil {
		fmt.Fprint(w, "Rejected")
	}
	conn, err := socket.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		mlogger.Println("Error establishing connection: ", err)
		return
	}
	connection := socket.Connection{
		User:       user,
		Connection: conn,
	}
	socket.UserConnections[user.ID] = connection
	for {
		messageType, p, err := connection.Connection.ReadMessage()
		if err != nil {
			log.Println(err)
			continue
		}
		if start == 0 {
			fmt.Println(string(p))
			start++
			continue
		}
		if messageType == websocket.TextMessage {
			msg := new(socket.MessageReader)
			err = json.Unmarshal(p, msg)
			if err != nil {
				log.Println("Error unmarshalling message: ", err)
				continue
			}
			fmt.Println(msg)
		}
	}
}
