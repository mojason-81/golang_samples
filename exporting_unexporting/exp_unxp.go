// From Go In Action - Manning Publications
// Sample program to show how the program can't access
// an unexported identifier from another package.
package main

import(
	"fmt"
	"github.com/goinaction/code/chapter5/listing68/counters"
)

// application entry point
func main() {
	// Create a variable of the unexported type and initialize
	// the value to 10.
	// counter := counters.alertCounter(10)

	// The above fails with - cannot refer to unexported name counters.alertCounter
	//                      - undefined: counters.alertCounter
	// NOTE: the "failing" sample needs to import from "listing64/counters"
	// The below New function exists in "listing68/counters"

	// Create a variable of the unexported type using the exported
	// New function from the package counters.
	counter := counters.New(10)

	fmt.Printf("Counter: %d\n", counter)
}
