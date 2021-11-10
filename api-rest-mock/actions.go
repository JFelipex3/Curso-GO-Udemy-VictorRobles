package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var movies = Movies{
	Movie{"Gladiator ", 2000, "Ridley Scott"},
	Movie{"El rey león", 1994, "Rob Minkoff, Roger Allers"},
	Movie{"Regreso al futuro", 1985, "Robert Zemeckis"},
	Movie{"El pianista", 2002, "Roman Polanski"},
	Movie{"Interstellar", 2014, "Christopher Nolan"},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movies)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	fmt.Fprintf(w, "Has cargado la película numero %s", movie_id)
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	log.Println(movie_data)
	movies = append(movies, movie_data)
}
