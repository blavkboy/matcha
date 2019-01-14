package socket

import (
	"fmt"
	"strings"

	"github.com/blavkboy/matcha/mlogger"
	"github.com/blavkboy/matcha/models"
	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

//Idea here is to use the connections map to keep track of all the users who are connected to the server.
//In the event that a user is online, a new connection should be established, the User's ID and a pointer
//to his/her connection is stored in the map. Should a user disconnect then we delete the connection.

//Connection will represent the way that connections are dealt with.
//When instanciated the models.User struct will be used to identify
//the user who's browser is sending a message
type Connection struct {
	User       *models.User
	Connection *websocket.Conn
}

type ProfileForm struct {
	Fname       string   `json:"fname"`
	Lname       string   `json:"lname"`
	Uname       string   `json:"uname"`
	Email       string   `json:"email"`
	Gender      string   `json:"gender"`
	Orientation string   `json:"orientation"`
	Interests   []string `json:"interests"`
}

type MessageReader struct {
	Type        string         `json:"type"`
	CommandType string         `json:"commandType"`
	Component   string         `json:"component"`
	Command     string         `json:"command"`
	Message     models.Message `json:"message"`
	Pform       ProfileForm    `json:"pform"`
}

var UserConnections = make(map[bson.ObjectId]Connection)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleMessage(msg *MessageReader) {
	if strings.Compare("message", string(msg.Type)) == 0 {
		message := new(models.Message)
		fmt.Println(message)
		return
	}
}

func (msg *MessageReader) HandleCommand(res *string, u *models.User) {
	mlogger := mlogger.GetInstance()
	if msg.CommandType == "profile" {
		if msg.Pform.Lname != "" {
			u.Lname = msg.Pform.Lname
		}
		if msg.Pform.Fname != "" {
			u.Fname = msg.Pform.Fname
		}
		if msg.Pform.Gender != "Select" {
			u.Sex = msg.Pform.Gender
		}
		if msg.Pform.Orientation != "Select" {
			u.Profile.Orientation = msg.Pform.Orientation
		}
		if len(msg.Pform.Interests) > 0 && msg.Pform.Interests[0] != "" {
			u.Profile.Interests = msg.Pform.Interests
		}
	} else if msg.CommandType == "propic" {
		u.Profile.Propic = msg.Command
	}
	err := u.UpdateUser()
	if err != nil {
		fmt.Println("Error updating user: ", err)
		mlogger.Println("Error updating user: ", err)
		*res = "{\"status\": false}"
	} else {
		fmt.Println("Success")
		mlogger.Println("Successully updated user_no: ", u.ID)
		*res = "{\"status\": true}"
	}
}

func (msg *MessageReader) EvalMsg(res *string, user *models.User, connection *Connection) {
	mlogger := mlogger.GetInstance()
	*res = ""
	switch msg.Type {
	case "command":
		msg.HandleCommand(res, user)
		err := connection.Connection.WriteMessage(websocket.TextMessage, []byte(*res))
		if err != nil {
			mlogger.Println("Error writing message to connection: ", err)
			return
		}
	}
}
