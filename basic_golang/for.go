package main

import "fmt"

func LoopingFor() {
	//for mirip foreach
	//print jika index angka genap
	title := "Golang"
	for index, char := range title {
		if (index % 2) == 0 {
			fmt.Println("Index ke : ", index, " hurufnya : ", string(char))
		}

		switch string(char) {
		case "a", "i", "u", "e", "o":
			fmt.Println("Huruf Vokal : ", string(char))
		default:
			fmt.Println("Non-Vokal : ", string(char))
		}
	}
}
