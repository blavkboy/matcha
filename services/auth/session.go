package auth

import (
	"net/http"

	"github.com/gmohlamo/matcha/mlogger"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(securecookie.GenerateRandomKey(32)))

//GetSession will get the session for the current request
//much safer than using jwt from localhost
func GetSession(r *http.Request) *sessions.Session {
	session, err := store.Get(r, "matcha-session")
	mlogger := mlogger.GetInstance()
	if err != nil {
		mlogger.Println(err)
		return nil
	}
	return session
}
