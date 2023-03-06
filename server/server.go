package main

import (
	"fmt"
	"net/http"
	"text/template"

	"groupietracker"
)

func main() {
	data := groupietracker.ApiArtiste()
	fs := http.FileServer(http.Dir("./front"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//initialisation des variables
		tmpl := template.Must(template.ParseFiles("template/index.html"))
		tmpl.Execute(w, data)
	})

	fmt.Println("server run on http://localhost:8080/ ")
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
}
