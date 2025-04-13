package db

import (
	"decoupled-data-storage/interfaces"
	"fmt"
)

// HashMapDB is a mock database implementation using Go's map.
// It's a simple in-memory key-value store that maps integer keys to IPerson records.
type HashMapDB map[int]interfaces.IPerson

// Save inserts a new IPerson record into the HashMapDB.
// The record is stored with an automatically assigned integer key.
func (h HashMapDB) Save(p interfaces.IPerson) {
	h[len(h)] = p
	fmt.Println("Successfully added", p.GetFirst(), p.GetLast())
}

// Retrieve searches for an IPerson record by its ID.
// If found, it returns the person; otherwise, it returns nil.
func (h HashMapDB) Retrieve(id string) interfaces.IPerson {
	for _, person := range h {
		if person.GetID() == id {
			return person
		}
	}
	return nil
}

// Delete removes an IPerson record from the HashMapDB by its ID.
// If the person is found, they are deleted from the database.
func (h HashMapDB) Delete(id string) {
	for i, person := range h {
		if person.GetID() == id {
			delete(h, i)
			fmt.Println("Removing", person.GetFirst(), person.GetLast(), "(", person.GetID(), ") from DB")
			return
		}
	}
}

// Show prints all stored IPerson records in the HashMapDB.
// Each record is displayed with its first name, last name, ID, and age.
func (h HashMapDB) Show() {
	for _, person := range h {
		fmt.Println(person)
	}
	fmt.Println()
}
