package groupietracker

import (
	"fmt"
	"net/http"
	"text/template"
)

func Server() {
	data := ApiArtiste()
	fs := http.FileServer(http.Dir("../front/"))
	http.Handle("/front/", http.StripPrefix("/front/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//initialisation des variables
		tmpl := template.Must(template.ParseFiles("../template/index.html"))
		tmpl.Execute(w, data)
	})

	fmt.Println("server run on http://localhost:8080/")

	http.ListenAndServe(":8080", nil)
}
