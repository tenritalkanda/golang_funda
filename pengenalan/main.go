package main

import (
	"fmt"
	"pengenalan/calculation"
)

func main() {
	fmt.Println("Hello, World!")

	result := calculation.Add(1, 2)

	fmt.Println(result)

	result = calculation.Multiply(2, 3)

	fmt.Println(result)
}
