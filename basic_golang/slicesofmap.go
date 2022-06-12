package main

import "fmt"

func ContohSliceOfMap() {
	//slice of map
	students := []map[string]string{
		{"name": "Budi", "age": "20"},
		{"name": "Joni", "age": "21"},
		{"name": "Joko", "age": "22"},
	}

	for _, value := range students {
		fmt.Println(value)
		fmt.Println(value["name"])
	}
}
