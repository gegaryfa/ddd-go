package mongo

import (
	"context"
	"time"

	"github.com/gegaryfa/tavern/domain/customer"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	db        *mongo.Database
	customers *mongo.Collection
}

// mongoCustomer is an internal type that is used to store a CustomerAggregate
// we make an internal struct for this to avoid coupling this mongo implementation to the customers aggregate.
// Mongo uses bson so we add tags for that
type mongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

func NewFromCustomer(c customer.Customer) mongoCustomer {
	return mongoCustomer{
		ID:   c.GetID(),
		Name: c.GetName(),
	}
}

// ToAggregate converts into a aggregate.Customer
// this could validate all values present etc
func (m mongoCustomer) ToAggregate() customer.Customer {
	c := customer.Customer{}

	c.SetID(m.ID)
	c.SetName(m.Name)

	return c

}

// New creates a new mongodb repository
func New(ctx context.Context, connectionString string) (*Repository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	// Find Metabot DB
	db := client.Database("ddd")
	customers := db.Collection("customers")

	return &Repository{
		db:        db,
		customers: customers,
	}, nil
}

func (r *Repository) Get(id uuid.UUID) (customer.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := r.customers.FindOne(ctx, bson.M{"id": id})

	var c mongoCustomer
	err := result.Decode(&c)
	if err != nil {
		return customer.Customer{}, err
	}
	// Convert to aggregate
	return c.ToAggregate(), nil
}

func (r *Repository) Add(c customer.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	internal := NewFromCustomer(c)
	_, err := r.customers.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(c customer.Customer) error {
	panic("to implement")
}
