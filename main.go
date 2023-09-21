package main
import (
	  "fmt"
	  "log"
	  "net/http"
)
func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}
func formHandler(w http.ResponseWriter, r *http.Request){
	if err:= r.ParseForm(); err!=nil{
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "post request successful")
	name:=r.FormValue("name")
	address:=r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}
func main(){
	fileserver:=http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver) // path, function (fileserver returns value of built-in type Handler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	fmt.Printf("Starting server at port 5500\n")
	if err:=http.ListenAndServe(":5500", nil); err!=nil{
		log.Fatal(err)
	}
}