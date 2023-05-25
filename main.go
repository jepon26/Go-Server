package main

import (
	"fmt"
	"go/parser"
	"log"
	"net/http"
)


func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Println(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprint(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("adress")
	fmt.Println(w, "name = %s\n", name)
	fmt.Println(w, "Adress = %s\n", address)
	
}




func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "hello") // Use fmt.Fprint to write "hello" to the ResponseWriter
}





func main (){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 3000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
		}
}

