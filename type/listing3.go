package main

import "fmt"

type notifier interface {
	notifiy()
}

type user2 struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Println("Sending user email to %s<%s>\n", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notifiy()
}

func main() {
	u := user2{"Bill", "bill@email.com"}

	sendNotification(&u)
}
