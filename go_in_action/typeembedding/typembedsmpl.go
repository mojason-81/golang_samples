// From Go In Action - Manning Publications
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
	user  // Embedding Type
	level string
}


// application entry point
func main() {
	// Create an admin user.
	ad := admin{
		user: user{
			name: "john smith",
			email: "john@example.com",
		},
		level: "super",
	}

	// We can access the inner type's method directly,
	ad.user.notify()

	// The inner type's method is promoted.
	ad.notify()

	// Or we can call the sendNotification() function directly
	// passing in a pointer to the admin
	// The embedded inner type's implementation of the interface
	// is "promoted" to the outer type.
	sendNotification(&ad)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}
