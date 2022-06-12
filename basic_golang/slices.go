package main

import "fmt"

func ContohSlice() {
	// slice adalah array yang dapat diubah ukurannya
	gamingConsole := []string{"Playstation", "Xbox", "Nintendo", "Wii", "WiiU"}

	// menambahkan elemen baru ke dalam slice
	gamingConsole = append(gamingConsole, "Switch")

	for _, value := range gamingConsole {
		fmt.Println(value)
	}

	fmt.Println("\nmengubah ukuran slice : ")
	// mengubah ukuran slice
	gamingConsole = gamingConsole[:2]

	for _, value := range gamingConsole {
		fmt.Println(value)
	}
}
