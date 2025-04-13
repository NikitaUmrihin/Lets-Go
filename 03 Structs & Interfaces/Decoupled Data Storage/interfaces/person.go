package interfaces

// 	IPerson defines the interface for a person entity in the system.
// 	It provides methods for retrieving personal details and displaying a greeting.
type IPerson interface {
	// SayHello prints a greeting message from the person.
	SayHello()

	// GetFirst returns the first name of the person.
	GetFirst() string

	// GetLast returns the last name of the person.
	GetLast() string

	// GetID returns a unique identifier associated with the person.
	GetID() string

	// GetAge returns the agse of the person.
	GetAge() int

	String() string
}
