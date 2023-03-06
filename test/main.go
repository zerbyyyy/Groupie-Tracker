package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	// structure de votre réponse JSON
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"Image"`
	CreationDate int      `json:"CreationDate"`
	FirstAlbum   string   `json:"FirstAlbum"`
	Member       []string `json:"Member"`
}

func main() {
	// appeler l'API
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// lire la réponse JSON
	var data Response
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	// exporter la réponse dans un serveur
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	fmt.Println("server run on http://localhost:8080/ ")
	http.ListenAndServe(":8080", nil)
}
