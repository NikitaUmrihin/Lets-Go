package main

import (
	"decoupled-data-storage/db"
	"decoupled-data-storage/interfaces"
	"decoupled-data-storage/structs"
	"fmt"
)

func main() {
	// 	Initialize people array
	ppl := [7]interfaces.IPerson{}
	ppl[0] = structs.Person{
		First: "jimmy",
		Last:  "jones",
		ID:    "654879874",
		Age:   25,
	}

	ppl[1] = structs.Student{
		Person: structs.Person{
			First: "chilly",
			Last:  "willy",
			ID:    "681248725",
			Age:   69,
		},
		Major: "Chemistry",
		GPA:   85.0,
	}

	ppl[2] = structs.Student{
		Person: structs.Person{
			First: "mellow",
			Last:  "mike",
			ID:    "87987954652",
			Age:   21,
		},
		Major: "Philosophy",
		GPA:   79.5,
	}

	ppl[3] = structs.Person{
		First: "bobby",
		Last:  "bobbs",
		ID:    "651279924",
		Age:   32,
	}

	ppl[4] = structs.Person{
		First: "robbie",
		Last:  "wright",
		ID:    "645644004",
		Age:   55,
	}

	ppl[5] = structs.Student{
		Person: structs.Person{
			First: "lazy",
			Last:  "liza",
			ID:    "79874651",
			Age:   29,
		},
		Major: "History",
		GPA:   68.4,
	}

	ppl[6] = structs.Person{
		First: "tommy",
		Last:  "thompson",
		ID:    "125326754",
		Age:   48,
	}

	//  Introduction round
	fmt.Println("________")
	for i := 0; i < len(ppl); i++ {
		ppl[i].SayHello()
	}
	fmt.Println("________")

	//  Set DB accessor
	//              = &PDblListDB   , if you want to use a double linked list
	//              = &HashMapDB    , if you want to use a has map
	dba := &db.HashMapDB{}

	// 	Initialize database
	dbService := structs.NewDBService(dba)

	// 	Get accessor concrete type
	accessor_type := fmt.Sprintf("%T", dba)

	// 	Print database type
	if accessor_type == "*db.HashMapDB" {
		fmt.Println("TESTING HASH MAP DATABASE  :")
	}
	if accessor_type == "*db.PDblListDB" {
		fmt.Println("TESTING DOUBLE LINKED LIST DATABASE  :")
	}

	// 	Insert people into the database
	for i := 0; i < len(ppl); i++ {
		dbService.Put(ppl[i])
	}
	fmt.Println()

	// 	Show the entire database
	dbService.Print()
	fmt.Println()

	// 	Remove people from database
	dbService.Pull("79874651")
	dbService.Print()

	dbService.Pull("651279924")
	dbService.Print()

	dbService.Pull("654879874")
	dbService.Print()

	dbService.Pull("125326754")
	dbService.Print()

	//	Add another person
	dbService.Put(ppl[0])
	dbService.Print()
}
