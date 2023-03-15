package groupietracker

import (
	"fmt"
	"net/http"
	"text/template"
)

type API struct {
	Artist   []Artiste
	Relation []Relation
}

func Server() {
	data := ApiArtiste()
	dataFinalRecup := ApiRelation()
	//fmt.Println(dataArtiste[0].Id)
	fs := http.FileServer(http.Dir("../front/"))
	http.Handle("/front/", http.StripPrefix("/front/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//initialisation des variables
		tmpl := template.Must(template.ParseFiles("../template/index.html"))
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/location", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("../template/location.html"))
		tmpl.Execute(w, dataFinalRecup)
	})

	// fmt.Println(dataFinalRecup)
	fmt.Println("server run on http://localhost:8080/")

	http.ListenAndServe(":8080", nil)
}
