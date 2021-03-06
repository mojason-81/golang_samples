// From Go In Action - Manning Publications
// Sample program to show how to write a simple version of curl
// using the io.Reader and io.Writer interface support.
package main

import(
	"fmt"
	"io"
	"net/http"
	"os"
)

// init is called before main.
func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./example2 <url>")
		os.Exit(-1)
	}
}

// main is the entry point for the application.
func main() {
	//Get a response from the web server.
	response, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// Copies from the Body to Stdout.
	io.Copy(os.Stdout, response.Body)
	if err := response.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
