package groupietracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Artiste struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"Image"`
	CreationDate int      `json:"CreationDate"`
	FirstAlbum   string   `json:"FirstAlbum"`
	Member       []string `json:"Member"`
}

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, _ := http.NewRequest("GET", url, nil)
	var p Artiste
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	err := json.Unmarshal([]byte(url), &p)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Artiste: ", p)
}

func RecupInfoArtiste(w http.ResponseWriter, r *http.Request) {
	info := &Artiste{}
	json.NewEncoder(w).Encode(info)

}
