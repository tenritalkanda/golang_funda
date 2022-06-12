package main

import "fmt"

func TestArray() {

	intArray := [5]int{1, 2, 3, 4, 5}
	stringArray := [5]string{"a", "b", "c", "d", "e"}

	undefinedLengthArrayOfColour := [...]string{
		"red",
		"green",
		"blue",
	}

	for _, value := range intArray {
		fmt.Println(value)
	}

	for _, value := range stringArray {
		fmt.Println(value)
	}

	for _, value := range undefinedLengthArrayOfColour {
		fmt.Println(value)
	}
}
