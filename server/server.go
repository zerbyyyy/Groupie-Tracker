package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./front"))
	fmt.Println("server run on http://localhost:8080/ ")
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	info := api.Artiste
	json.NewEncoder(w).Encode(info)
}
