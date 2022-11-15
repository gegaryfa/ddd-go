package tavern

import (
	"testing"

	"github.com/gegaryfa/tavern/domain/customer"
	"github.com/gegaryfa/tavern/domain/product"
	"github.com/gegaryfa/tavern/services/order"
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

func TestTavern_Order(t *testing.T) {
	// Create OrderService
	products := initProducts(t)

	os, err := order.NewOrderService(
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	cust, err := customer.NewCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	err = os.Customers.Add(cust)
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Execute Order
	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}

func Test_MongoTavern(t *testing.T) {
	t.Skip("this will fail when no mongo is running")
	// Create OrderService
	products := initProducts(t)

	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository("mongodb://localhost:27017"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	uid, err := os.AddCustomer("Percy")
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}

	// Execute Order
	err = tavern.Order(uid, order)
	if err != nil {
		t.Error(err)
	}

}
