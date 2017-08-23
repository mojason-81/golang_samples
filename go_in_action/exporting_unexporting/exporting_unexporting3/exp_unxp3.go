// From Go In Action - Manning Publications
// Sample program to show how unexported fields from an exported
// struct type can't be accessed directly.
package main

import (
	"fmt"

	"github.com/goinaction/code/chapter5/listing74/entities"
)

// app entry point
func main() {
	// Create a value of type User from the entities package.
	u := entities.Admin{
		Rights: 10,
	}

	// Set the exported fields from the unexported inner type.
	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", u)
}
