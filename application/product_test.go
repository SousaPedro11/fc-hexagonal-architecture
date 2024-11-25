package application_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/sousapedro11/fc-arquitetura-hexagonal/application"
	"github.com/stretchr/testify/assert"
)

func TestProductIsValid(t *testing.T) {
	productId := uuid.NewString()
	productName := "Product 1"
	tests := []struct {
		name     string
		product  application.ProductInterface
		expected bool
		err      bool
	}{
		{
			name: "Valid Product",
			product: &application.Product{
				Price:  10,
				Id:     productId,
				Name:   productName,
				Status: application.ENABLED,
			},
			expected: true,
			err:      false,
		},
		{
			name: "Invalid Product Status",
			product: &application.Product{
				Price:  10,
				Id:     productId,
				Name:   productName,
				Status: "invalid",
			},
			expected: false,
			err:      true,
		},
		{
			name: "Invalid Product Name",
			product: &application.Product{
				Price:  10,
				Id:     productId,
				Name:   "",
				Status: application.ENABLED,
			},
			expected: false,
			err:      true,
		},
		{
			name: "Invalid Product Id",
			product: &application.Product{
				Price:  10,
				Id:     "invalid",
				Name:   productName,
				Status: application.ENABLED,
			},
			expected: false,
			err:      true,
		},
		{
			name: "Invalid Product Price",
			product: &application.Product{
				Price:  -10,
				Id:     productId,
				Name:   productName,
				Status: application.ENABLED,
			},
			expected: false,
			err:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.product.IsValid()
			if tt.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestProductEnable(t *testing.T) {
	productId := uuid.NewString()
	productName := "Product 2"
	tests := []struct {
		name     string
		product  *application.Product
		expected string
		err      bool
	}{
		{
			name: "Enabled Successful",
			product: &application.Product{
				Price:  10,
				Name:   productName,
				Id:     productId,
				Status: application.DISABLED,
			},
			expected: application.ENABLED,
			err:      false,
		},
		{
			name: "Enabled Failed - Price equals zero",
			product: &application.Product{
				Name:   productName,
				Id:     productId,
				Status: application.DISABLED,
			},
			expected: application.DISABLED,
			err:      true,
		},
		{
			name: "Enabled Failed - Already Enabled",
			product: &application.Product{
				Price:  10,
				Name:   productName,
				Id:     productId,
				Status: application.ENABLED,
			},
			expected: application.ENABLED,
			err:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.product.Enable()
			if tt.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.product.GetStatus(), tt.expected)
		})
	}
}

func TestProductDisable(t *testing.T) {
	productId := uuid.NewString()
	productName := "Product 3"
	tests := []struct {
		name     string
		product  *application.Product
		expected string
		err      bool
	}{
		{
			name: "Disable Successful",
			product: &application.Product{
				Name:   productName,
				Id:     productId,
				Status: application.ENABLED,
				Price:  0,
			},
			expected: application.DISABLED,
			err:      false,
		},
		{
			name: "Disable Failed - Price greater than zero",
			product: &application.Product{
				Name:   productName,
				Id:     productId,
				Status: application.ENABLED,
				Price:  10,
			},
			expected: application.ENABLED,
			err:      true,
		},
		{
			name: "Disable Failed - Already Disabled",
			product: &application.Product{
				Name:   productName,
				Id:     productId,
				Status: application.DISABLED,
				Price:  0,
			},
			expected: application.DISABLED,
			err:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.product.Disable()
			if tt.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.product.GetStatus(), tt.expected)
		})
	}
}

func TestProductChangePrice(t *testing.T) {
	productId := uuid.NewString()
	productName := "Product 4"
	tests := []struct {
		name     string
		product  *application.Product
		price    float64
		expected float64
		err      bool
	}{
		{
			name: "Change Price Successful - Price greater than zero",
			product: &application.Product{
				Name:   productName,
				Id:     productId,
				Status: application.ENABLED,
				Price:  10,
			},
			price:    20,
			expected: 20,
			err:      false,
		},
		{
			name: "Change Price Successful - Price equals zero",
			product: &application.Product{
				Name:   productName,
				Id:     productId,
				Status: application.ENABLED,
				Price:  10,
			},
			price:    0,
			expected: 0,
			err:      false,
		},
		{
			name: "Change Price Failed - Price less than zero",
			product: &application.Product{
				Name:   productName,
				Id:     productId,
				Status: application.ENABLED,
				Price:  10,
			},
			price:    -10,
			expected: 10,
			err:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.product.ChangePrice(tt.price)
			if tt.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expected, tt.product.GetPrice())
		})
	}
}

func TestProductGetters(t *testing.T) {
	productId := uuid.NewString()
	productName := "Product 5"
	product := &application.Product{
		Name:   productName,
		Id:     productId,
		Status: application.ENABLED,
		Price:  10,
	}

	assert.Equal(t, productId, product.GetId())
	assert.Equal(t, productName, product.GetName())
	assert.Equal(t, application.ENABLED, product.GetStatus())
	assert.Equal(t, 10.0, product.GetPrice())
}

func TestProductConstructor(t *testing.T) {
	productName := "Product 6"
	product := application.NewProduct(productName, 10)

	assert.NotEmpty(t, product.GetId())
	assert.Equal(t, productName, product.GetName())
	assert.Equal(t, application.DISABLED, product.GetStatus())
	assert.Equal(t, 10.0, product.GetPrice())
}
