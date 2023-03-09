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
	req, errNR := http.NewRequest("GET", url, nil)
	if errNR != nil {
		fmt.Println("Error Request preparation:", errNR)
		return data
	}
	res, errDC := http.DefaultClient.Do(req)
	if errDC != nil {
		fmt.Println("Error Request send :", errDC)
		return data
	}
	defer res.Body.Close()
	body, ioutil := ioutil.ReadAll(res.Body)
	if ioutil != nil {
		fmt.Println("Error ioutil :", ioutil)
		return data
	}
	//fmt.Println(string(body))
	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error", err)
		return data
	}
	//fmt.Println("Artiste: ", data)
	return data
}

/*func RecupInfoArtiste(w http.ResponseWriter, r *http.Request) {
	info := &Artiste{}
	json.NewEncoder(w).Encode(info)

}
*/
