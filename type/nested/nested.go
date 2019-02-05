package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email %s<%s>\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

func main() {
	ad := admin{
		user: user{
			name:  "John smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	ad.user.notify()

	ad.notify()
}
