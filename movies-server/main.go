package main
import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"math/rand"
	"github.com/gorilla/mux"
	"strconv"
)
// Movie Struct (Model)
type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"` //pointer to Director
}
type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}
// Init movies var as a slice Movie struct
var movies []Movie


func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json") //content type is json (informing the broser)
	json.NewEncoder(w).Encode(movies) //to encode slice into json and send it back to the browser
}
func deleteMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	params:=mux.Varse(r)
	for index, item:= range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}
func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= mux.Varse(r)
	for index, item := range movies{
		if item.id == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	
}
func main(){
	r:=mux.NewRouter() //create a new router
	// Mock Data - @todo - implement DB
	movies = append(movies, Movie{ID: "1", Isbn: "448743", Title:"Movie One", Director: &Director{Firstname:"John", Lastname:"Doe"}})
	movies = append(movies, Movie{ID:"2", Isbn:"786876", Title:"Movie 2", Director: &Director{Firstname:"Steve", Lastname:"Smith"}})
	movies = append(movies, Movie{ID:"3", Isbn:"987987", Title:"Movie 3", Director: &Director{Firstname:"Jane", Lastname:"Doe"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Printf("Starting server at port 5500\n")
	log.Fatal(http.ListenAndServe(":5500",r))
}