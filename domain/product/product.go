package product

import (
	"errors"

	"github.com/gegaryfa/tavern"
	"github.com/google/uuid"
)

var (
	// ErrMissingValues is returned when a product is created without a name or description
	ErrMissingValues = errors.New("missing values")
)

type Product struct {
	item     *tavern.Item
	price    float64
	quantity int
}

func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValues
	}

	return Product{
		item: &tavern.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *tavern.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}
