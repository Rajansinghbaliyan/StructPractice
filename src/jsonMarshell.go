package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var myJson = `
[
	{
	"first_name": "Harry",
	"last_name": "Potter"
	},

	{
	"first_name": "James",
	"last_name": "Bond"
	}
]
`

func main() {
	readJson()
	writeJson()
}

func readJson() {
	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println(err)
	}
	log.Println(unmarshalled)

}

func writeJson() {
	harryBosch := Person{
		FirstName: "Harry",
		LastName:  "Bosch",
	}

	data, err := json.MarshalIndent(&harryBosch, "", "   ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(data))
}
