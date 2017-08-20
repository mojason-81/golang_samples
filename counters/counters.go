// From Go In Action - Manning Publications
// Sample program to show how the program can't access an
// unexported identifier from another package.
package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
