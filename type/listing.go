package main

import "fmt"

type user_list struct {
	name  string
	email string
}

// 实现定义结构体的方法
func (u user_list) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

// 实现定义结构体的方法(使用指针的方式)
func (u *user_list) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user_list{"Bill", "bill@email.com"}
	bill.notify()

	lisa := &user_list{"Lisa", "lisa@email.com"}
	lisa.notify()

	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	lisa.changeEmail("lisa@comcast.com")
	lisa.notify()
}
