package auth

import (
	"github.com/gmohlamo/matcha/mlogger"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte(securecookie.GenerateRandomKey(32)))

func GetSession(r *http.Request) *sessions.Session {
	session, err := store.Get(r, "matcha-session")
	mlogger := mlogger.GetInstance()
	if err != nil {
		mlogger.Println(err)
		return nil
	}
	return session
}
