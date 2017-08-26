// This sample program demonstrates how to decode a JSON string.
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Contact represents our JSON string.
type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

// JSON contains a sample string to unmarshal.
var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
	}
}`

func main() {
	// Unmarshal the JSON string into our new variable.
	var c Contact
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(c)

	// Unmarshal the JSON string into our map variable
	var co map[string]interface{}
	err = json.Unmarshal([]byte(JSON), &co)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println("Name: ", co["name"])
	fmt.Println("Title: ", co["title"])
	fmt.Println("Contact")
	fmt.Println("\tH: ", co["contact"].(map[string]interface{})["home"])
	fmt.Println("\tC: ", co["contact"].(map[string]interface{})["cell"])
}
