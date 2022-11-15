package memory

import (
	"fmt"
	"sync"

	"github.com/gegaryfa/tavern/domain/customer"
	"github.com/google/uuid"
)

type repository struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

func New() *repository {
	return &repository{customers: make(map[uuid.UUID]customer.Customer)}
}

func (r repository) Get(uuid uuid.UUID) (customer.Customer, error) {
	if customer, ok := r.customers[uuid]; ok {
		return customer, nil
	}
	return customer.Customer{}, customer.ErrCustomerNotFound
}

func (r repository) Add(c customer.Customer) error {
	if _, ok := r.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}

	r.Mutex.Lock()
	r.customers[c.GetID()] = c
	r.Mutex.Unlock()

	return nil
}

func (r repository) Update(c customer.Customer) error {
	// Make sure Customer is in the repository
	if _, ok := r.customers[c.GetID()]; !ok {
		return customer.ErrCustomerNotFound
	}
	r.Lock()
	r.customers[c.GetID()] = c
	r.Unlock()
	return nil
}
