package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	FirstName string `json:"first"`
	LastName string `json:"last"`
	Hair string `json:"hair"`
	Dog bool `json:"has_dog"`
}

func main(){

	// write json into struct

	myJson := `
	[
		{
			"first": "clark",
			"last": "kent",
			"hair": "black",
			"has_dog": true
		},
		{
			"first": "dave",
			"last": "peter",
			"hair": "blue",
			"has_dog": false
		}
	]
	`

		var unmarshalled []Person

		err := json.Unmarshal([]byte(myJson), &unmarshalled) // Need to convert from string to slice of bytes, and second param interface you put slice of bytes into. Put it into reference to unmarshalled (slice)
		if err != nil {
			fmt.Println("Error unmarshalling json", err)
		}

		fmt.Printf("unmarshalled %v \n", unmarshalled) // read json into a struct

		//write json from a struct

		var mySlice []Person

		var m1 Person
		m1.FirstName = "doggo"
		m1.LastName = "red"
		m1.Hair = "blue"
		m1.Dog = false

		mySlice = append(mySlice, m1)

		var m2 Person
		m2.FirstName = "dianna"
		m2.LastName = "prince"
		m2.Hair = "green"
		m2.Dog = false

		mySlice = append(mySlice, m2)

		// Now slice with two entries. Convert this into JSON

		newJson, err := json.MarshalIndent(mySlice, "", "   ") //Indent formats nicely. Prefix is second param, and indent spacing is third param.

		if err != nil {
			fmt.Println("Error marshalling", err)
		}

		fmt.Println(string(newJson))
}
