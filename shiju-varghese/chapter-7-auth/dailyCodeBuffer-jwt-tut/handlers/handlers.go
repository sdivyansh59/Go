package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

// creting users and giving then username and pass
// in real life we will get then from DB
var users = map[string]string {
	"user1" : "password1",
	"user2" : "password2",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}


func Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return 
	}

	expectedPassword, ok := users[credentials.Username]

	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute* 5)

	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// now with the help of jwtKey and claim obj 
	// we wil generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString , err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	// set cookie
	http.SetCookie(w,
		&http.Cookie{
			Name : "token",
			Value : tokenString,
			Expires: expirationTime,
		},
	)


}

func Home(w http.ResponseWriter, r *http.Request) {
	cookie ,err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return 
		}
		w.WriteHeader(http.StatusBadRequest)
		return 
	}

	tokenStr := cookie.Value

	claim := &Claims{}

	tkn , err := jwt.ParseWithClaims(tokenStr, claim,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return 
		}
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return 
	}

	w.Write([]byte(fmt.Sprintf("Hello %s ",claim.Username)))
}



func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie ,err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return 
		}
		w.WriteHeader(http.StatusBadRequest)
		return 
	}

	tokenStr := cookie.Value

	claim := &Claims{}

	tkn , err := jwt.ParseWithClaims(tokenStr, claim,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return 
		}
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return 
	}

	expirationTime := time.Now().Add(time.Minute * 5)
	claim.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	tokenString , err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	http.SetCookie(w,
		&http.Cookie{
			Name : "Refresh_token",
			Value : tokenString,
			Expires: expirationTime,
		},
	)



}