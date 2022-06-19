package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

type Gamer struct {
	Name  string
	Games []string
}

func (games *Gamer) AddGame(game string) {
	games.Games = append(games.Games, game)
}

func main() {
	numberA := 10
	numberB := 20

	fmt.Println("Before change:"+" numberA:", numberA, " numberB:", numberB)
	change(&numberA, numberB)
	fmt.Println("After change:"+" numberA:", numberA, " numberB:", numberB)

	studentA := Student{"John", 20}
	fmt.Println("Before graduate:"+" studentA:", studentA)
	graduate(&studentA)
	fmt.Println("After graduate:"+" studentA:", studentA)

	gamer := Gamer{
		Name: "John",
	}

	gamer.AddGame("ABCD")
	gamer.AddGame("1234")
	gamer.AddGame("XYVZ")

	for index, game := range gamer.Games {
		fmt.Println(index, ""+game)
	}
}

func change(old *int, new int) {
	*old = new
}

func graduate(student *Student) {
	student.Age = 24
	student.Name = student.Name + " P.hD"
}
