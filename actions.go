package main

import "fmt"
import "net/http"
import "log"
import "github.com/gorilla/mux"
import "encoding/json"
import "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"


//MongoDB Connect
func getSession() *mgo.Session{
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil{
		panic(err)
	}

	return session
}
//MongoDB Connect

//variable glogal para MongoDB
var collection = getSession().DB("curso_go").C("movies")

var movies = Movies{
	Movie{"Harry Potter 1",2001,"Desconocido"},
	Movie{"Harry Potter 2",2002,"Desconocido"},
	Movie{"Harry Potter 3",2003,"Desconocido"},
	Movie{"Harry Potter 4",2004,"Cuarón"},
}

//Funcion no aplicada, se determinó no usar el responses
/*
func responseMovies (w http.ResponseWriter, status int, results [] Movie){
	w.Header().Set("Content Type", "aplication/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(movie_data)
}
*/


func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con Go con Gorilla/mux")
}

func Contacto(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Contacto con Router Gorilla/mux")
}

func MovieList(w http.ResponseWriter, r *http.Request){

	var results []Movie
	
	//err := collection.Find(nil).All(&results) //del primero al actual
	err := collection.Find(nil).Sort("-_id").All(&results) //orden del actual al primero

	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Println("Resultados: ", results)
	}

	/*
	movies := Movies{
		Movie{"Harry Potter 1",2001,"Desconocido"},
		Movie{"Harry Potter 2",2002,"Desconocido"},
		Movie{"Harry Potter 3",2003,"Desconocido"},
		Movie{"Harry Potter 4",2004,"Cuarón"},
	}
	*/
	//fmt.Fprintf(w, "Listado de peliculas")
	//json.NewEncoder(w).Encode(movies)
	w.Header().Set("Content Type", "aplication/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

func MovieShow(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	movie_id := params["id"]
	/**/
	//hay que convertir el parametro id en un hexadecimal para poder usarlo en el Json binario 

	if !bson.IsObjectIdHex(movie_id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)
	fmt.Println(movie_id)
	fmt.Println(oid)
	results := Movie{} //objeto movie vacio movie.go
	err := collection.FindId(oid).One(&results)

	fmt.Println(results)

	if err != nil{
		w.WriteHeader(404)
		return
	}else{
		w.Header().Set("Content Type", "aplication/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(results)
	}
	/**/

	//fmt.Fprintf(w, "Show Pelicula ")
	//fmt.Fprintf(w,"Has cargado la pelicula numero %s", movie_id)


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

	//log.Println(movie_data)

	//Insert en MongoDB
	//session := getSession() //Cambio por la variable global "arriba"
	//session.DB("curso_go").C("movies").Insert(movie_data) //Cambio por la variable global "arriba"
	err = collection.Insert(movie_data)

	if err != nil {
		w.WriteHeader(500)
		return 
	}
	//Insert en MongoDB

	//movies = append(movies, movie_data) //comentado por la implementacion de la BD

	w.Header().Set("Content Type", "aplication/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(movie_data)
}



func MovieUpdate(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	movie_id := params["id"]

	
	/**/
	//hay que convertir el parametro id en un hexadecimal para poder usarlo en el Json binario 
	
	if !bson.IsObjectIdHex(movie_id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil{
		panic(err)
		w.WriteHeader(500)
		return
	}

	defer r.Body.Close()
	
	document := bson.M{"_id": oid}
	change := bson.M{"$set": movie_data}

	err = collection.Update(document,change)

	if err != nil{
		w.WriteHeader(404)
		return
	}
	
	w.Header().Set("Content Type", "aplication/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(movie_data)
	
	/**/

	//fmt.Fprintf(w, "Show Pelicula ")
	//fmt.Fprintf(w,"Has cargado la pelicula numero %s", movie_id)


}

type Message struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

//metodo para unir el struct Message con una funcion y servir de intercambio de valores "Extra"
func (this *Message) setStatus (data string){
	this.Status = data
}

func (this *Message) setMessage (data string){
	this.Message = data
}

func MovieRemove(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	movie_id := params["id"]
	/**/
	//hay que convertir el parametro id en un hexadecimal para poder usarlo en el Json binario 

	if !bson.IsObjectIdHex(movie_id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)


	err := collection.RemoveId(oid)

	if err != nil{
		w.WriteHeader(404)
		return
	}


	//results := Message{"success","La pelicula con ID " + movie_id + " ha sido borrada correctamente"} // cambio por los metodos de intercambio del struct Message "extra"
	message := new(Message)
	message.setStatus("success")
	message.setMessage("La pelicula con ID " + movie_id + " ha sido borrada correctamente")

	results := message

	w.Header().Set("Content Type", "aplication/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)


}