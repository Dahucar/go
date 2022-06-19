package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Definbicion de tipos de datos
type Movie struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Director    *Director `json:"duration"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Almacenamiento temporal
var movies []Movie

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{ID: "01", Title: "Example 01", Description: "Awesome movie", Director: &Director{Firstname: "Dan", Lastname: "H"}})
	movies = append(movies, Movie{ID: "02", Title: "Example 02", Description: "Awesome movie", Director: &Director{Firstname: "Dan", Lastname: "H"}})
	movies = append(movies, Movie{ID: "03", Title: "Example 03", Description: "Awesome movie", Director: &Director{Firstname: "Dan", Lastname: "H"}})
	movies = append(movies, Movie{ID: "04", Title: "Example 04", Description: "Awesome movie", Director: &Director{Firstname: "Dan", Lastname: "H"}})

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movie", addMovie).Methods("POST")
	router.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")
	router.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8881", router))
}

// Operaciones del servidor.
func addMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	fmt.Println(params)
	for index, movie := range movies {
		fmt.Print(index)
		fmt.Print(movie)
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	fmt.Println(params)

	for index, movie := range movies {
		fmt.Print(index)
		fmt.Print(movie)
		if movie.ID == params["id"] {
			// Eliminar el elemento de array
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	fmt.Println(params)
	for index, movie := range movies {
		fmt.Print(index)
		fmt.Print(movie)
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+index:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}
