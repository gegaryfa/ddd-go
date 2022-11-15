// Package customer holds aggregate that combines many entities into a full object
package customer

import (
	"errors"

	"github.com/gegaryfa/tavern"
	"github.com/google/uuid"
)

var (
	// ErrInvalidName is returned when the name is not valid in the NewCustomer factory
	ErrInvalidName = errors.New("a customer has to have an valid name")
)

// Customer is an aggregate that combines all entities needed to represent a customer
// Note: I set all the entities as pointers, this is because an entity can change state and I want that to reflect
// across all instances of the runtime that has access to it. The value objects are held as nonpointers though since they cannot change state.
type Customer struct {
	// person is the root entity of a customer
	// which means the person.ID is the main identifier for this aggregate
	person *tavern.Person
	// a customer can hold many products
	products []*tavern.Item
	// a customer can perform many transactions
	transactions []tavern.Transaction
}

// NewCustomer is a factory to create a new Customer aggregate
// It will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidName
	}

	// Create a new person and generate ID
	person := &tavern.Person{
		ID:   uuid.New(),
		Name: name,
	}

	// Create a customer object and initialize all the values to avoid nil pointer exceptions
	return Customer{
		person:       person,
		products:     make([]*tavern.Item, 0),
		transactions: make([]tavern.Transaction, 0),
	}, nil
}

func (c Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}
	c.person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}
	c.person.Name = name
}

func (c Customer) GetName() string {
	return c.person.Name
}
