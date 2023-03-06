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

func ApiArtiste() []Artiste {
	var data []Artiste

	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, ioutil := ioutil.ReadAll(res.Body)
	if ioutil != nil {
		fmt.Println("Error ioutil :", ioutil)
	}
	//fmt.Println(string(body))
	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Artiste: ", data[0].CreationDate)
	return data
}

/*func RecupInfoArtiste(w http.ResponseWriter, r *http.Request) {
	info := &Artiste{}
	json.NewEncoder(w).Encode(info)

}
*/
