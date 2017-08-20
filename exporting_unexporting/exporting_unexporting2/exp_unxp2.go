// From Go In Action - Manning Publications
// Sample program to show how unexported fields from an exported
// struct type can't be accessed directly.
package main

import (
	"fmt"

	"github.com/goinaction/code/chapter5/listing71/entities"
)

// app entry point
func main() {
	// Create a value of type User from the entities package.
	u := entities.User{
		Name:  "Bill",
		email: "bill@email.com",
	}

	// the above fails with - unknown entities.User field 'email' in struct literal

	fmt.Printf("User: %v\n", u)
}
