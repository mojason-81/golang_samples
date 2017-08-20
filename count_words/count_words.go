// Declares a single function that
// takes a quoted string as a path to a file,
// counts the words in that file, and outputs
// to the user the number of words.
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/goinaction/code/chapter3/words"
)

// The main function that makes calls to other packages
// to read the file and print the response.
func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("There are %d words in %s. \n", count, filename)
}
