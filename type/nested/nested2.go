package main

import (
	"fmt"
)

// 实现嵌套结构中的接口覆盖的实现
type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

// admin 自己实现notify，这样在ad.notify的时候不会调user的实现
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

func sendNotfication(n notifier) {
	n.notify()
}

func main() {
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	sendNotfication(&ad)

	ad.notify()

	ad.user.notify()
}
