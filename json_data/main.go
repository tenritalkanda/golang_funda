package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dogs"`
}

func main() {
	myJson := `
	[
		{
			"first_name":"John",
			"last_name":"Doe",
			"hair_color":"blue",
			"has_dogs":true
		},
		{
			"first_name":"Jane",
			"last_name":"Doe",
			"hair_color":"red",
			"has_dogs":false
			}
	]`

	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		panic(err)
	}

	for _, person := range unmarshalled {
		fmt.Printf("%s %s has %s hair and has %t dogs \n", person.FirstName, person.LastName, person.HairColor, person.HasDog)
	}

	mySlice := []Person{}
	m1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		HairColor: "blue",
		HasDog:    true,
	}

	m2 := Person{
		FirstName: "Jane",
		LastName:  "Doe",
		HairColor: "red",
		HasDog:    false,
	}

	mySlice = append(mySlice, m1)
	mySlice = append(mySlice, m2)

	// newJson, err := json.MarshalIndent(mySlice, "", "   ")
	newJson, err := json.Marshal(mySlice)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(newJson))

}
