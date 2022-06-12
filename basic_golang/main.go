package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//for mirip foreach
	title := "Golang"
	for index, char := range title {
		fmt.Println("Index ke : ", index, " hurufnya : ", string(char))
	}

}
