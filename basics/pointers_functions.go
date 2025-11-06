package main

import "fmt"

type User struct {
	name  string
	age   int
	admin bool
}

func (u *User) Create() {
	u.name = "James"
	u.age = 41
	u.admin = true
}

func Print(u *User, i int) {
	u.name = "Bond"
	u.age += i
	u.admin = false
}

func Show(u User) User {
	u.name = "Smith"
	u.age += 10
	u.admin = true
	return u
}

func callUser() {
	u := User{}
	u.Create()

	fmt.Println(u) // Print the initial User from Create {James 41 true}

	Print(&u, 7)

	fmt.Println(u) // Print the modified User from Print {Bond 48 false}

	fmt.Println(Show(u)) // Print the returned User from Show {Smith 58 true}
}

func main() {
	callUser()
}
