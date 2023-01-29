package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// using asymmetric crypto/RSa keys
//location of the files used for signing and verification

const (
	privKeyPath = "Keys/app.rsa" // openssl generate key
	pubKeyPath = "Keys/app.rsa.pub" // 
)

// verify key and sign key
var (
	verifyKey, signKey []byte
)

//struct user for parsing login credentials
type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// read the key files before starting http handlers
func init() {
	var err error 

	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return 
	}

	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading public key")
		return 
	}
}

// read the login credentials , check them and created JWT token
func loginHandler(w http.ResponseWriter, r * http.Request){
	var user User
	// deode into \User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w,"Error in request body ")
		return 
	}
}

//validate user credentials
if user.UserName != "shivjuvar" && user.Password != "pass" {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Println(w,"Wrong info")
	return
}

func main() {
	fmt.Println("Hello")



}