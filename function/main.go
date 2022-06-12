package main

import (
	"errors"
	"fmt"
)

func main() {
	tambah := add(10, 20)
	fmt.Println(tambah)

	luas, keliling := keliling(10, 20)
	fmt.Println(luas)
	fmt.Println(keliling)

	scores := []int{10, 5, 8, 9, 7}
	total := sum_array(scores)
	fmt.Println(total)

	// result, err := calculate(10, 2, "+")
	// result, err := calculate(10, 2, "-")
	// result, err := calculate(10, 2, "*")
	// result, err := calculate(10, 2, "/")
	result, err := calculate(10, 2, "=")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

func add(a, b int) int {
	return a + b
}

func keliling(panjang, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	return luas, keliling
}

//func dengan pre-defined return value
func keliling2(panjang, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	return luas, keliling
}

func sum_array(scores []int) (total int) {
	for _, v := range scores {
		total += v
	}
	return total
}

func calculate(a, b int, operator string) (int, error) {
	var result int
	var errorResult error

	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		errorResult = errors.New("operator tidak dikenali")
	}

	return result, errorResult
}
