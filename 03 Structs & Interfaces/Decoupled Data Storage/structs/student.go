package structs

import "fmt"

//	Student has same fields as Person (first name, last name , id, and age)
// 	+ major and GPA
type Student struct {
	Person //  Cool trick -> no name for variable person (lets us use fields directly)
	Major  string
	GPA    float64
}

// 	SayHello implementation
func (s Student) SayHello() {
	fmt.Println("Hi, I'm", s.First, s.Last, ", I study", s.Major)
}

// 	First name getter
func (s Student) GetFirst() string {
	return s.First
}

// 	Last name getter
func (s Student) GetLast() string {
	return s.Last
}

// 	ID getter
func (s Student) GetID() string {
	return s.ID
}

// 	Age getter
func (s Student) GetAge() int {
	return s.Age
}

// 	Major getter
func (s Student) GetMajor() string {
	return s.Major
}

// 	GPA getter
func (s Student) GetGPA() float64 {
	return s.GPA
}

//
func (s Student) String() string {
	return fmt.Sprintf("ID: %s => %s %s is %d years old, I study %s and have %f GPA", s.ID, s.First, s.Last, s.Age, s.Major, s.GPA)
}
