package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nylo-andry/movies-service/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", handlers.AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies", handlers.CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", handlers.UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", handlers.DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/{id}", handlers.FindMovieEndpoint).Methods("GET")

	log.Println("Routes registered")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
