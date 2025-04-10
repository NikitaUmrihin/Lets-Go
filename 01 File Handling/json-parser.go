package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Define a User struct
type User struct {

	// Struct tags allow the program to correctly map the incoming JSON data to the struct fields
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// String() lets a type define its custom print format used by fmt.
func (u User) String() string {
	return fmt.Sprintf("Hello! My name is %s, I am %d years old.\t contact me at: %s\n", u.Name, u.Age, u.Email)
}

// Let's Go!
func main() {

	// Read the JSON file
	data, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Declare a slice of Users to hold the parsed data
	var users []User

	// Parse the JSON data into the users slice
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Print the data in a structured format
	for _, user := range users {
		fmt.Println(user)
	}
}
