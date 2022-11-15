package order

import (
	"testing"

	"github.com/gegaryfa/tavern/domain/customer"
	"github.com/gegaryfa/tavern/domain/product"
	"github.com/google/uuid"
)

func initProducts(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		t.Error(err)
	}
	peenuts, err := product.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := product.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	products := []product.Product{
		beer, peenuts, wine,
	}
	return products
}

func TestOrder_NewOrderService(t *testing.T) {
	// Create a few Products to insert into in memory repo
	products := initProducts(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Error(err)
	}

	// Add Customer
	customer, err := customer.NewCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	err = os.Customers.Add(customer)
	if err != nil {
		t.Error(err)
	}

	// Perform Order for one beer
	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(customer.GetID(), order)

	if err != nil {
		t.Error(err)
	}

}
