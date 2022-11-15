package memory

import (
	"testing"

	"github.com/gegaryfa/tavern/domain/customer"
	"github.com/google/uuid"
)

func Test_Add(t *testing.T) {
	type testCase struct {
		name        string
		customer    string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Add Customer",
			customer:    "Percy",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := repository{
				customers: map[uuid.UUID]customer.Customer{},
			}

			c, err := customer.NewCustomer(tc.customer)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(c)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(c.GetID())
			if err != nil {
				t.Fatal(err)
			}
			if found.GetID() != c.GetID() {
				t.Errorf("Expected %v, got %v", c.GetID(), found.GetID())
			}
		})
	}
}

func Test_Get(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	// Create a fake customer to add to repository
	cust, err := customer.NewCustomer("Percy")
	if err != nil {
		t.Fatal(err)
	}
	id := cust.GetID()
	// Create the repo to use, and add some test Data to it for testing
	// Skip Factory for this
	repo := repository{
		customers: map[uuid.UUID]customer.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:        "No Customer By ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: customer.ErrCustomerNotFound,
		}, {
			name:        "Customer By ID",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_, err := repo.Get(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func Test_Update(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
		newName     string
	}

	// Create a fake customer to add to repository
	newCustomer, err := customer.NewCustomer("Percy")
	if err != nil {
		t.Fatal(err)
	}
	id := newCustomer.GetID()
	// Create the repo to use, and add some test Data to it for testing
	// Skip Factory for this
	repo := repository{
		customers: map[uuid.UUID]customer.Customer{
			id: newCustomer,
		},
	}

	testCases := []testCase{
		{
			name:        "Customer does not exist",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			newName:     "George",
			expectedErr: customer.ErrCustomerNotFound,
		}, {
			name:        "Update customer name",
			id:          id,
			newName:     "George",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, err := customer.NewCustomer("TestCustomer")
			if err != nil {
				t.Fatal(err)
			}

			c.SetID(tc.id)
			c.SetName(tc.newName)

			err = repo.Update(c)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
