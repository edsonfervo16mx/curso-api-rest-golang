package main

import "fmt"
import "net/http"
import "log"
import "github.com/gorilla/mux"

func main(){
	fmt.Println("Server start 5050")
	//metodo nativo, sustituido por el metodo de gorilla/mux
	/*
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hola mundo desde mi servidor web con Go")
	})
	*/

	//metodo mux
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/contacto", Contacto)

	router.HandleFunc("/peliculas/", MovieList)
	router.HandleFunc("/peliculas/{id}", MovieShow)

	//para el nativo
	//server := http.ListenAndServe(":5050",nil)

	//para el gorilla/mux
	server := http.ListenAndServe(":5050",router)

	log.Fatal(server)	
}


func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con Go con Gorilla/mux")
}

func Contacto(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Contacto con Router Gorilla/mux")
}

func MovieList(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Listado de peliculas")
}

func MovieShow(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	movie_id := params["id"]


	fmt.Fprintf(w, "Show Pelicula")
	fmt.Fprintf(w,"Has cargado la pelicula numero %s", movie_id)
}