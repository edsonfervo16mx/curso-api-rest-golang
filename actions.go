package main

import "fmt"
import "net/http"
import "log"
import "github.com/gorilla/mux"
import "encoding/json"

var movies = Movies{
	Movie{"Harry Potter 1",2001,"Desconocido"},
	Movie{"Harry Potter 2",2002,"Desconocido"},
	Movie{"Harry Potter 3",2003,"Desconocido"},
	Movie{"Harry Potter 4",2004,"Cuarón"},
}

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con Go con Gorilla/mux")
}

func Contacto(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Contacto con Router Gorilla/mux")
}

func MovieList(w http.ResponseWriter, r *http.Request){
	/*
	movies := Movies{
		Movie{"Harry Potter 1",2001,"Desconocido"},
		Movie{"Harry Potter 2",2002,"Desconocido"},
		Movie{"Harry Potter 3",2003,"Desconocido"},
		Movie{"Harry Potter 4",2004,"Cuarón"},
	}
	*/
	//fmt.Fprintf(w, "Listado de peliculas")
	json.NewEncoder(w).Encode(movies)
}

func MovieShow(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	movie_id := params["id"]


	fmt.Fprintf(w, "Show Pelicula")
	fmt.Fprintf(w,"Has cargado la pelicula numero %s", movie_id)
}

//Ejemplo de peticion POST
func MovieAdd(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if(err != nil){
		panic(err)
	}

	defer r.Body.Close()

	log.Println(movie_data)
	movies = append(movies, movie_data)
}