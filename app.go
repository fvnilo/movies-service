package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"github.com/nylo-andry/movies-service/config"
	"github.com/nylo-andry/movies-service/models"
	"github.com/nylo-andry/movies-service/repository"
)

var dbConfig = config.Config{}
var movieRepository = repository.MovieRepository{}

// AllMoviesEndPoint returns a collection of all the movies in the database
func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	movies, err := movieRepository.FindAll()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, movies)
}

// FindMovieEndpoint returns a single movie or null that has the provided id
func FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := movieRepository.FindByID(params["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}

	respondWithJSON(w, http.StatusOK, movie)
}

// CreateMovieEndPoint creates a movies with the provided information.
// The request must have a "name", a "cover_image" and a "description".
func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie models.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	movie.ID = bson.NewObjectId()

	if err := movieRepository.Insert(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, movie)
}

// UpdateMovieEndPoint updates a movies with the provided information.
// The request must have a "name", a "cover_image" and a "description".
func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie models.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := movieRepository.Update(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteMovieEndPoint deletes a movies with a corresponding id.
func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie models.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := movieRepository.Delete(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	dbConfig.Read()

	movieRepository.Server = dbConfig.Server
	movieRepository.Database = dbConfig.Database

	log.Printf("Initializing server connection on [%s] [%s]", dbConfig.Server, dbConfig.Database)

	movieRepository.Connect()
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/{id}", FindMovieEndpoint).Methods("GET")

	log.Println("Routes registered")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
