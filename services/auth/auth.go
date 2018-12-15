package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/blavkboy/matcha/mlogger"
	"github.com/blavkboy/matcha/models"
	"github.com/blavkboy/matcha/routing"
	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("The `jig is up")

//NewToken produces a new token
func NewToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.Compare(r.Method, "POST") == 0 {
			var user models.User
			var compare *models.User
			json.NewDecoder(r.Body).Decode(&user)
			fmt.Println("Recieved data: ", user)
			// Create a new token object, specifying signing method and the claims
			// you would like it to contain.
			compare = models.FindUser("username", user.Username)
			if strings.Compare(compare.Password, user.Password) != 0 {
				routing.HandleRoot(w, r)
				next(w, r)
			}
			fmt.Println("user pw: ", user.Password)
			fmt.Println("compare pw: ", compare.Password)
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"user":    compare,
				"created": time.Now(),
				"exp":     time.Now().AddDate(0, 0, 7),
			})
			// Sign and get the complete encoded token as a string using the secret
			tokenString, err := token.SignedString(mySigningKey)
			if err != nil {
				logger := mlogger.GetInstance()
				logger.Println("Error: ", err)
				return
			}

			exp := time.Now().Add(time.Hour * (24 * 7))
			cookie := http.Cookie{Name: authToken, Value: tokenString, Expires: exp}
			http.SetCookie(w, &cookie)
		}
	}
}

func ConfirmUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := mlogger.GetInstance()
		cookie, err := r.Cookie(authToken)
		if err != nil {
			failedAuth(w, r)
			return
		}
		tokenString := cookie.Value
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				logger.Println("Error: Unexpected signing method: ", token.Header["alg"])
				return nil, fmt.Errorf("Error: Unexpected signing method: %v", token.Header["alg"])
			}

			return mySigningKey, nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			logger.Println(claims["user"])
		} else {
			logger.Println("Error: failed to authorize user")
			failedAuth(w, r)
			return
		}
		next(w, r)
	}
}

func failedAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Authorization Failed")
}

const authToken = "AuthToken17286983217313"
