package main

import "fmt"
import "net/http"
import "log"
//import "github.com/gorilla/mux"
//import "encoding/json"

func main(){
	fmt.Println("Server start 5050")
	//metodo nativo, sustituido por el metodo de gorilla/mux
	/*
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hola mundo desde mi servidor web con Go")
	})
	*/

	//metodo mux

	/*COMENTADO POR QUE SE HIZO ARCHIVO ROUTER*/
	/*
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/contacto", Contacto)

	router.HandleFunc("/peliculas/", MovieList)
	router.HandleFunc("/peliculas/{id}", MovieShow)

	*/
	router := NewRouter()
	//


	//para el nativo
	//server := http.ListenAndServe(":5050",nil)

	//para el gorilla/mux
	server := http.ListenAndServe(":5050",router)

	log.Fatal(server)	
}

