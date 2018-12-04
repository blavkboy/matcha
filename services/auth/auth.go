package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/blavkboy/matcha/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("The `jig is up")

//NewToken produces a new token
func NewToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.Compare(r.Method, "POST") == 0 {
			var user models.User
			json.NewDecoder(r.Body).Decode(&user)
			// Create a new token object, specifying signing method and the claims
			// you would like it to contain.
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"user":    user,
				"created": time.Now(),
				"exp":     time.Now().AddDate(0, 0, 7),
			})
			models.Users = append(models.Users, user)
			r.URL.Path = r.URL.Path + "/" + user.ID

			// Sign and get the complete encoded token as a string using the secret
			tokenString, err := token.SignedString(mySigningKey)

			fmt.Println(tokenString, err)
			fmt.Println(token)
			exp := time.Now().Add(time.Hour * 48)
			cookie := http.Cookie{Name: authToken, Value: tokenString, Expires: exp}
			http.SetCookie(w, &cookie)
		}
		confirmUser(w, r, next)
	}
}

func confirmUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	cookie, err := r.Cookie(authToken)
	if err != nil {
		failedAuth(w, r)
		return
	}
	tokenString := cookie.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return mySigningKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Fprint(w, "\n", claims)
	} else {
		fmt.Fprint(w, "Error, failed to authorize user")
	}
	next(w, r)
}

func failedAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Authorization Failed")
}

const authToken = "AuthToken"
