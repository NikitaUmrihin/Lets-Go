package interfaces

// 	Accessor is how to store or retrieve a user in our database.
// 	When retrieving a user that doesn't exist zero value is returned
type Accessor interface { //  An accessor can :
	Save(p IPerson)             //  Save a user to database
	Retrieve(id string) IPerson //  Get a user by index in database
	Delete(id string)           //  Delete a user by index in database
	Show()                      //  Show the entire database
}
