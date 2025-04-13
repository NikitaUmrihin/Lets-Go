package structs

import "fmt"

//	Person has first name, last name , id, and age
type Person struct {
	First string
	Last  string
	ID    string
	Age   int
}

// 	SayHello implementation
func (p Person) SayHello() {
	fmt.Println("Hi, I'm", p.First, p.Last)
}

// 	First name getter
func (p Person) GetFirst() string {
	return p.First
}

// 	Last name getter
func (p Person) GetLast() string {
	return p.Last
}

// 	ID getter
func (p Person) GetID() string {
	return p.ID
}

// 	Age getter
func (p Person) GetAge() int {
	return p.Age
}

func (p Person) String() string {
	return fmt.Sprintf("ID: %s => %s %s is %d years old", p.ID, p.First, p.Last, p.Age)
}
