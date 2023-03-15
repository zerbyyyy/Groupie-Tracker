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
	Member       []string `json:"Members"`
}

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type PreRelation struct {
	Index []Relation `json:"index"`
}

type RelationName struct {
	Name           []string
	RelationUpdate Relation
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
	// fmt.Println(string(body))
	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error", err)
		return data
	}
	// fmt.Println("Artiste: ", data[0])
	return data
}

func ApiRelation() []Relation {
	var dataRelation []Relation
	var dataPreRelation PreRelation

	url := "https://groupietrackers.herokuapp.com/api/relation"
	req, ErrorNewRequest := http.NewRequest("GET", url, nil)
	if ErrorNewRequest != nil {
		fmt.Println("Error Request preparation: ", ErrorNewRequest)
		return dataRelation
	}
	res, ErrorDC := http.DefaultClient.Do(req)
	if ErrorDC != nil {
		fmt.Println("Error request send: ", ErrorDC)
		return dataRelation
	}
	defer res.Body.Close()
	body, ioutil := ioutil.ReadAll(res.Body)
	if ioutil != nil {
		fmt.Println("Error ioutil :", ioutil)
		return dataRelation
	}
	// fmt.Println(string(body))
	err := json.Unmarshal(body, &dataPreRelation)
	if err != nil {
		fmt.Println("Error Unmarshall 2: ", err)
		return dataRelation
	}
	for _, relation := range dataPreRelation.Index {
		dataRelation = append(dataRelation, relation)
	}

	return dataRelation
}

/*
func RecupName() []string {
	var test = ApiRelation()
	var test2 Artiste
	var resultName []string
	url := ""
	for i := 1; i < len(test)+1; i++ {
		url = "https://groupietrackers.herokuapp.com/api/artists/" + strconv.Itoa(i)
		req, ErrorNewRequest := http.NewRequest("GET", url, nil)
		if ErrorNewRequest != nil {
			fmt.Println("Error Request preparation : ", ErrorNewRequest)
			return resultName
		}
		res, ErrorDC := http.DefaultClient.Do(req)
		if ErrorDC != nil {
			fmt.Println("Error Request Send : ", ErrorDC)
			return resultName
		}
		defer res.Body.Close()
		body, ioutil := ioutil.ReadAll(res.Body)
		if ioutil != nil {
			fmt.Println("Error ioutil : ", ioutil)
			return resultName
		}
		err := json.Unmarshal(body, &test2)
		if err != nil {
			fmt.Println("Error Unmarshall : ", err)
			return resultName
		}
		// fmt.Println(test2.Name)
		resultName = append(resultName, test2.Name)
	}
	fmt.Println(resultName)
	return resultName
}


func LastRecupName() RelationName {
	var RecupTable RelationName
	Stock := RecupName()
	StockRelation := ApiRelation()
	fmt.Print("StockRelation : ")
	fmt.Println(StockRelation[1])
	for i := 0; i < len(Stock); i++ {
		RecupTable.Name = append(RecupTable.Name, Stock[i])
		//RecupTable.RelationUpdate = append(RecupTable.RelationUpdate, StockRelation[i])
	}
	return RecupTable
}
*/
