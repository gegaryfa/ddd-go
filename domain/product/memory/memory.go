package memory

import (
	"sync"

	"github.com/gegaryfa/tavern/domain/product"
	"github.com/google/uuid"
)

type Repository struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}

func New() *Repository {
	return &Repository{
		products: make(map[uuid.UUID]product.Product),
	}
}

func (r Repository) GetAll() ([]product.Product, error) {
	var products []product.Product
	for _, product := range r.products {
		products = append(products, product)
	}

	return products, nil
}

func (r Repository) GetByID(id uuid.UUID) (product.Product, error) {
	if product, ok := r.products[id]; ok {
		return product, nil
	}
	return product.Product{}, product.ErrProductNotFound

}

func (r Repository) Add(newProduct product.Product) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.products[newProduct.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}

	r.products[newProduct.GetID()] = newProduct

	return nil
}

func (r Repository) Update(upprod product.Product) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.products[upprod.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	r.products[upprod.GetID()] = upprod
	return nil
}

func (r Repository) Delete(id uuid.UUID) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.products[id]; !ok {
		return product.ErrProductNotFound
	}
	delete(r.products, id)
	return nil
}
