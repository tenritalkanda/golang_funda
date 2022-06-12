package main

import "fmt"

func Scores() {
	var avgscore float64
	var sumScores int
	var goodScores []int

	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	for _, value := range scores {
		sumScores += value
		if value >= 90 {
			goodScores = append(goodScores, value)
		}
	}

	avgscore = float64(sumScores) / float64(len(scores))
	fmt.Println(avgscore)

	fmt.Println(goodScores)
}
