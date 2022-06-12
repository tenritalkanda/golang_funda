package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func (user User) displayUser() string {
	return fmt.Sprintf("Name : %s %s, Email : %s\n", user.FirstName, user.LastName, user.Email)
}

func (awaw Group) displayGroup() {
	fmt.Printf("Name : %s\n", awaw.Name)
	fmt.Printf("Member Count : %d\n", len(awaw.Users))

	// for _, user := range group.Users {
	// 	fmt.Println(user.FirstName)
	// }
}

func main() {
	user := User{1, "John", "Doe", "test@gmail.com", true}
	user2 := User{2, "Sist", "Doe", "test2@gmail.com", true}

	result := user.displayUser()
	fmt.Println(result)

	users := []User{user, user2}
	group := Group{"Gamer", user, users, true}

	// displayGroup(group)

	group.displayGroup()

}

// func displayGroup(group Group) {
// 	fmt.Printf("Name : %s\n", group.Name)
// 	fmt.Printf("Member Count : %d\n", len(group.Users))

// 	for _, user := range group.Users {
// 		fmt.Println(user.FirstName)
// 	}

// }
