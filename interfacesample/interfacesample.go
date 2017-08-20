// Sample program to show how to use an interface in Go.
package main

import "fmt"

// notifier is an interface that defines notification
// type behavior.
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a pointer receiver.
func(u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

// admin defines an admin in th eprogram.
type admin struct {
	name string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

// application entry point
func main() {
	// Create a value of type User and send a notification.
	u := user{"Bill", "bill@email.com"}
	sendNotification(&u)

	// Create an admin value and pass it to sendNotification.
	lisa := admin{"Lisa", "lisa@email.com"}
	sendNotification(&lisa)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}
