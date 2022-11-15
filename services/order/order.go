package order

import (
	"context"
	"log"

	"github.com/gegaryfa/tavern/domain/customer"
	"github.com/gegaryfa/tavern/domain/customer/memory"
	"github.com/gegaryfa/tavern/domain/customer/mongo"
	"github.com/gegaryfa/tavern/domain/product"
	prodmemory "github.com/gegaryfa/tavern/domain/product/memory"
	"github.com/google/uuid"
)

// OrderConfiguration is an alias for a function that will take in a pointer to an OrderService and modify it
type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	Customers customer.Repository
	Products  product.Repository
}

// See how we can take in a variable amount of OrderConfiguration in the factory method? It is a very neat way
// of allowing dynamic factories and allows the developer to configure the architecture, given that it is implemented.
// This trick is very good for unit tests, as you can replace certain parts in service with the wanted repository.
func NewOrderService(cfg ...OrderConfiguration) (*OrderService, error) {
	// Create the order service
	os := &OrderService{}
	// Apply all Configurations passed in
	for _, cfg := range cfg {
		// Pass the service into the configuration function
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithCustomerRepository(cr customer.Repository) OrderConfiguration {
	// return a function that matches the OrderConfiguration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(os *OrderService) error {
		os.Customers = cr
		return nil
	}
}

// WithMemoryCustomerRepository applies a prodmemory customer repository to the OrderService
func WithMemoryCustomerRepository() OrderConfiguration {
	// Create the prodmemory repo, if we needed parameters, such as connection strings they could be inputted here
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMongoCustomerRepository(connectionString string) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the mongo repo, if we needed parameters, such as connection strings they could be inputted here
		cr, err := mongo.New(context.Background(), connectionString)
		if err != nil {
			return err
		}
		os.Customers = cr
		return nil
	}
}

// WithMemoryProductRepository adds a in prodmemory product repo and adds all input Products
func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the prodmemory repo, if we needed parameters, such as connection strings they could be inputted here
		pr := prodmemory.New()

		// Add Items to repo
		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}
		os.Products = pr
		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	// get the customer
	c, err := o.Customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	//todo
	var products []product.Product
	var price float64
	for _, id := range productIDs {
		p, err := o.Products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		price += p.GetPrice()
	}

	// All Products exist in store, now we can create the order
	log.Printf("Customer: %s has ordered %d Products", c.GetID(), len(products))

	return price, nil
}

// AddCustomer will add a new customer
func (o OrderService) AddCustomer(name string) (uuid.UUID, error) {
	c, err := customer.NewCustomer(name)
	if err != nil {
		return uuid.Nil, err
	}

	err = o.Customers.Add(c)
	if err != nil {
		return uuid.Nil, err
	}

	return c.GetID(), nil

}
