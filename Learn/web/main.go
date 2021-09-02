package main

import (
	"log"
	"net/http"
)

var (
	port = "8000"
)

func main() {
	http.HandleFunc("/", myweb)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func myweb(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
}
