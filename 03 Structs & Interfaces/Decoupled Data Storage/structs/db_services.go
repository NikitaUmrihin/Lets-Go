package structs

import (
	"decoupled-data-storage/interfaces"
	"fmt"
)

// DBService acts as a wrapper around a database accessor, providing
// high-level operations for managing IPerson records.
type DBService struct {
	A interfaces.Accessor
}

// NewDBService acts as a constructor -
// creates and returns a new DBService instance with the provided Accessor.
func NewDBService(a interfaces.Accessor) DBService {
	return DBService{
		A: a,
	}
}

// Get retrieves an IPerson record from the database by ID.
// If the person does not exist, it returns an error.
func (db DBService) Get(id string) (interfaces.IPerson, error) {
	p := db.A.Retrieve(id)
	if p.GetFirst() == "" && p.GetLast() == "" {
		return Person{}, fmt.Errorf("ERR: no person with id %s", id)
	}
	return p, nil
}

// Put saves a new IPerson record into the database using the configured Accessor.
func (db DBService) Put(p interfaces.IPerson) {
	db.A.Save(p)
}

// Pull removes an IPerson record from the database by ID.
func (db DBService) Pull(id string) {
	db.A.Delete(id)
}

// Print shows the entire database.
func (db DBService) Print() {
	fmt.Println()
	fmt.Println("ALL ENTRIES:")
	db.A.Show()
}
