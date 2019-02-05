package main

import "fmt"

type nofitier interface {
	notify()
}

func (u *user) notify() {
	fmt.Println("Sending user eamil to %s<%s>\n", u.name, u.email)
}

type admin struct {
	name  string
	email string
}

func (a *admin) notify() {
	fmt.Println("Sending admin email to %s<%s>\n", a.name, a.email)
}

func main() {
	bill := user2{"Bill", "bill@email.com"}
	sendNotification(&bill)

	lisa := admin("Lisa", "lisa@email.com")
	sendNotification(&lisa)
}

func sendNotification(n notifier) {
	n.notifiy()
}
