package main

import "fmt"

func main() {
	numberA := 5
	numberB := &numberA

	fmt.Println(numberA, numberB)
}
