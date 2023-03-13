package main

import (
	"fmt"
)

func main() {
	data := map[string]int{"Evan": 18, "Paul": 20, "Spouf": 15}

	for i := range data {
		fmt.Println("le personage s'apelle ", i, "est a ", data[i])
	}

}
