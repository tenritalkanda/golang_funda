package main

import "fmt"

func MyMap() {
	//map adalah array yg menggunakan key:value
	var mymap = map[string]int{}
	mymap["a"] = 1
	mymap["b"] = 2
	mymap["c"] = 3

	fmt.Println(mymap)

	//menghapus elemen map
	delete(mymap, "a")
	fmt.Println(mymap)

	//mencari elemen map
	value, exists := mymap["a"]
	fmt.Println(value, exists)

	yourmap := map[string]int{"ruby": 1, "python": 2, "php": 3}
	fmt.Println(yourmap)
	fmt.Println(yourmap["ruby"])
}
