package product

import (
	"reflect"
	"testing"

	"github.com/gegaryfa/tavern"
	"github.com/google/uuid"
)

func TestNewProduct(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		description string
		price       float64
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "should return error if name is empty",
			name:        "",
			expectedErr: ErrMissingValues,
		},
		{
			test:        "validvalues",
			name:        "test",
			description: "test",
			price:       1.0,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := NewProduct(tc.name, tc.description, tc.price)
			if err != tc.expectedErr {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}

func TestProduct_GetID(t *testing.T) {
	type fields struct {
		item     *tavern.Item
		price    float64
		quantity int
	}
	tests := []struct {
		name   string
		fields fields
		want   uuid.UUID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Product{
				item:     tt.fields.item,
				price:    tt.fields.price,
				quantity: tt.fields.quantity,
			}
			if got := p.GetID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProduct_GetItem(t *testing.T) {
	type fields struct {
		item     *tavern.Item
		price    float64
		quantity int
	}
	tests := []struct {
		name   string
		fields fields
		want   *tavern.Item
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Product{
				item:     tt.fields.item,
				price:    tt.fields.price,
				quantity: tt.fields.quantity,
			}
			if got := p.GetItem(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProduct_GetPrice(t *testing.T) {
	type fields struct {
		item     *tavern.Item
		price    float64
		quantity int
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Product{
				item:     tt.fields.item,
				price:    tt.fields.price,
				quantity: tt.fields.quantity,
			}
			if got := p.GetPrice(); got != tt.want {
				t.Errorf("GetPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
