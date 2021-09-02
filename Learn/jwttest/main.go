package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//custom claims openinfo//
type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Data string `json:"data"`
}

type token struct {
	Token string `json:"token"`
}

var jwtSecret = []byte("secret")

func Login(w http.ResponseWriter, r *http.Request) {
	var user UserCredentials
	err := json.NewDecoder(r.Body).Decode(&user)
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	router := gin.Default()

}
