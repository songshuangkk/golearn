package main

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type admin struct {
	person user
	level  string
}

func main() {
	var bill user

	fmt.Println(bill)

	lisa := user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}

	fmt.Println(lisa)

	lisa2 := user{"Lisa2", "lisa2@email.com", 123, true}

	fmt.Println(lisa2)

	fred := admin{
		person: lisa,
		level:  "2",
	}

	fmt.Println(fred)
}
