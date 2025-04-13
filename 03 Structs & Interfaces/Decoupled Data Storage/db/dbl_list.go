package db

import (
	"decoupled-data-storage/interfaces"
	"fmt"
)

// Node represents a single element in a doubly linked list-based database.
// It stores an IPerson data entry and pointers to the previous and next nodes.
type Node struct {
	Data interfaces.IPerson
	Next *Node
	Prev *Node
}

// PDblListDB is a mock database implementation using a doubly linked list.
// It maintains insertion order and is optimized for sequential access.
type PDblListDB struct {
	Head *Node //  first node
	Tail *Node //  last node
	Size int
}

// Append adds a new node to the end of the doubly linked list.
// If the list is empty, the new node becomes both the head and tail.
func (list *PDblListDB) Append(n *Node) {
	if list.Size == 0 {
		list.Tail = n
		list.Head = n
		list.Size++
		return
	}
	temp := list.Tail
	temp.Next = n
	n.Prev = temp
	list.Tail = n
	list.Size++
}

// Save inserts a new IPerson record into the doubly linked list.
// It wraps the IPerson instance in a Node and appends it to the list.
func (list *PDblListDB) Save(p interfaces.IPerson) {
	newNode := Node{Data: p}
	list.Append(&newNode)
	fmt.Println("Successfully added", p.GetFirst(), p.GetLast())
}

// Retrieve searches for an IPerson record by its ID in the linked list.
// If found, it returns the person; otherwise, it returns nil.
func (list *PDblListDB) Retrieve(id string) interfaces.IPerson {
	curr := list.Head
	for i := 0; i < list.Size && curr != nil; i++ {
		if curr.Data.GetID() == id {
			return curr.Data
		}
		curr = curr.Next
	}

	return nil
}

// Delete removes an IPerson record from the doubly linked list by its ID.
func (list *PDblListDB) Delete(id string) {
	curr := list.Head
	for i := 0; i < list.Size && curr != nil; i++ {

		if curr.Data.GetID() == id {
			fmt.Println("Removing", curr.Data.GetFirst(), curr.Data.GetLast(), "(", id, ") from DB")
			if curr.Prev != nil {
				curr.Prev.Next = curr.Next
			} else {
				list.Head = curr.Next
			}
			if curr.Next != nil {
				curr.Next.Prev = curr.Prev
			} else {
				list.Tail = curr.Prev
			}
			list.Size--
		}
		curr = curr.Next

	}
}

// Show prints all stored IPerson records in the linked list in order.
// Each record is displayed with its first name, last name, ID, and age.
func (list *PDblListDB) Show() {
	curr := list.Head
	for curr != nil {
		fmt.Println(curr.Data)
		curr = curr.Next
	}
	fmt.Println()
}
