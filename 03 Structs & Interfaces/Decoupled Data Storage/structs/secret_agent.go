package structs

import "fmt"

//	SecretAgent has same fields as Person (first name, last name , id, and age)
// 	plus a boolean field Ltk - representing license to kill
type SecretAgent struct {
	Person      //  Cool trick -> no name for variable person (lets us use fields directly)
	Ltk    bool //	License to Kill
}

// SayHello implementation
func (sa SecretAgent) SayHello() {
	if sa.Ltk {
		fmt.Println("You know too much, I'm gonna have to kill you now...", sa.Last)
	} else {
		fmt.Println("I'm", sa.Last, "-", sa.First, sa.Last)
	}
}

// 'License to kill' getter
func (sa SecretAgent) GetLicense() bool {
	return sa.Ltk
}
