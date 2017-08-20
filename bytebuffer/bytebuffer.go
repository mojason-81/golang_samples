// From Go In Action - Manning Publications
// Sample program to show how a bytes.Buffer can also be
// used with the io.Copy function.
package main

import(
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer

	// Write a string to the buffer.
	b.Write([]byte("Hello"))
	
	// Use Fprintef to concatenate a string to the Buffer.
	fmt.Fprintf(&b, "World!")
	fmt.Fprintf(&b, "\n")

	// Write the content of Buffer to stdout.
	io.Copy(os.Stdout, &b)
}
