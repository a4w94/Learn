package main

import (
	m "movie/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", m.AllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", m.FindlMovies).Methods("GET")
	r.HandleFunc("/movies", m.CreateMovies).Methods("POST")
	r.HandleFunc("/movies", m.UpdateMovies).Methods("PUT")
	r.HandleFunc("/movies", m.DeleteMovies).Methods("DELETE")

	http.ListenAndServe(":8080", r)

}
