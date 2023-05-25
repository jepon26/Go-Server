package main

import (
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name    string
	Address string
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	user := User{
		Name:    r.FormValue("name"),
		Address: r.FormValue("address"),
	}

	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, user)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/form", formHandler)
	http.Handle("/", http.FileServer(http.Dir("static")))

	log.Println("Starting server at port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}





