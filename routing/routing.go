package routing

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gmohlamo/matcha/mlogger"
	"github.com/gmohlamo/matcha/models"
	"github.com/gmohlamo/matcha/services/auth"
	"github.com/gmohlamo/matcha/services/validation"
	"github.com/gmohlamo/matcha/socket"
	"github.com/gmohlamo/matcha/views"
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

func HandleLikes(w http.ResponseWriter, r *http.Request) {
	session := auth.GetSession(r)
	user := models.FindUser("username", session.Values["user"])
	if user == nil {
		w.Write([]byte("{\"success\": false}"))
		return
	}
	decoder := json.NewDecoder(r.Body)
	var like models.Like
	err := decoder.Decode(&like)
	if err != nil {
		w.Write([]byte("{\"success\": false}"))
		return
	}
	if like.Uid == user.ID {
		if models.AddLike(&like) {
			w.Write([]byte("{\"success\": true}"))
			return
		}
	}
	w.Write([]byte("\"success\": false"))
}

func HandleMatches(w http.ResponseWriter, r *http.Request) {
	mlogger := mlogger.GetInstance()
	w.Header().Set("Content-Type", "application/json")
	if strings.Compare(r.Method, "GET") == 0 {
		mlogger.Println("Sending user: ")
		session := auth.GetSession(r)
		user := models.FindUser("username", session.Values["user"])
		if user == nil {
			w.Write([]byte("{\"success\": false}"))
			return
		}
		models.FindMatch(user, w)
	} else {
		mlogger.Println("Got the right method")
		w.Write([]byte("{\"success\": false}"))
	}
}

func checkReg(body models.User) bool {
	mlogger := mlogger.GetInstance()
	if len(body.Fname) == 0 {
		mlogger.Println("first name has nothing")
		return false
	} else if len(body.Lname) == 0 {
		mlogger.Println("last name has nothing")
		return false
	} else if validation.ValidEmail(body.Email) != true {
		mlogger.Panicln("invalid email")
		return false
	} else if len(body.Password) <= 7 {
		mlogger.Panicln("password is too short")
		return false
	}
	return true
}

//HandleUser can be used to either Register a new user account or it can be used
//to retrieve user information for the client requesting it. The latter requires
//a valid session cookie to work
func HandleUser(w http.ResponseWriter, r *http.Request) {
	mlogger := mlogger.GetInstance()
	if strings.Compare(r.Method, "POST") == 0 {
		//we process get request and return either the selected user or
		//all the users
		//Todo: refine search capabilities and make this more efficient
		var body models.User
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewDecoder(r.Body).Decode(&body)
		if checkReg(body) == false {
			w.Write([]byte("{\"success\": false}"))
		}
		pass, err := bcrypt.GenerateFromPassword([]byte(body.Password), models.Cost)
		if err != nil {
			fmt.Println("Error hashing password: ", err.Error())
			return
		}
		body.Password = string(pass)
		mlogger.Println("Saving user with password: ", string(body.Password))
		newBody := models.FindUser("email", body.Email)
		if newBody != nil {
			w.Write([]byte("{\"success\": false, \"field\": \"email\"}"))
			return
		}
		newBody = models.FindUser("username", body.Username)
		if newBody != nil {
			w.Write([]byte("{\"success\": false, \"field\": \"username\"}"))
			return
		}
		ret := models.NewUser(&body)
		if ret == nil {
			w.Write([]byte("{\"success\": false, \"cause\": \"failed\"}"))
			return
		}
		w.Write([]byte("{\"success\": true}"))
		mlogger.Println(time.Now())
		mlogger.Println(body)
	} else {
		session := auth.GetSession(r)
		if session == nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("{\"success\": false}"))
		} else {
			user := models.FindUser("ID", session.Values["user"])
			if user == nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("{\"success\": false}"))
			} else {
				user.Password = ""
				body, err := json.Marshal(user)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("{\"success\": false}"))
					return
				}
				w.Write([]byte("{\"success\": true, \"user\"" + string(body) + "}"))
			}
		}
	}
}

//HandleUsers will do something
func HandleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

//HandleCheck I forgot what it does
func HandleCheck(w http.ResponseWriter, r *http.Request) {
	if strings.Compare(r.Method, "GET") == 0 {
		fmt.Println("Got Here")
		fmt.Println(r.Header)
		w.Header().Set("Content-Type", "application/json")
		mlogger := mlogger.GetInstance()
		user := auth.GetCurrentUser(r)
		if user == nil {
			mlogger.Println("User does not exist for token: ", r.Header.Get("Authorization"))
			w.Write([]byte("{\"success\": false}"))
		} else {
			mlogger.Println("User no: " + user.ID + " has logged into the Matcha service.")
			w.Write([]byte("{\"success\": true}"))
		}
	}
}

//HandleHome was for the ego thing
func HandleHome(w http.ResponseWriter, r *http.Request) {
	views.RenderHome(w, auth.GetCurrentUser(r))
}

//HandleUpdate should process user account updates
func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	session := auth.GetSession(r)
	user := models.FindUser("ID", session.Values["user"])
	if user == nil {
		w.Write([]byte("{\"success\": false}"))
		return
	}
	fmt.Printf("Found the user %v\n", user)
	var updatedUser models.User
	json.NewDecoder(r.Body).Decode(&updatedUser)
	if user.CheckUpdate(updatedUser) == false {
		w.Write([]byte("{\"success\": false}"))
	} else {
		user.UpdateDiff(updatedUser)
		user.Password = ""
		json.NewEncoder(w).Encode(user)
	}
}

//SocketConn should do socket things...
func SocketConn(w http.ResponseWriter, r *http.Request) {
	var start int = 0
	mlogger := mlogger.GetInstance()
	session := auth.GetSession(r)
	user := models.FindUser("username", session.Values["user"])
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
		fmt.Println("Got connection")
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
			msg.EvalMsg(user, &connection)
		}
	}
}
