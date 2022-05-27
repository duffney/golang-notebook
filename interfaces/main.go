package main

import (
	"fmt"
)

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
	name  string
	email string
}

func (a admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	u := user{"Bill", "bill@email.com"}
	sendNotification(&u) // pass a pointer of type user
	// sendNotification(u)
	// err: “Cannot use u (type user) as type notifier in argument
	//       to sendNotification: user does not implement notifier”
	a := admin{"John", "john@email.com"}
	sendNotification(a)  // pass a value of type admin
	sendNotification(&a) // pass a pointer of type admin
}
