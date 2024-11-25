package application_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sousapedro11/fc-arquitetura-hexagonal/application"
	"github.com/sousapedro11/fc-arquitetura-hexagonal/application/mock"
	"github.com/stretchr/testify/assert"
)

func TestProductServiceGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPersistence := mock.NewMockProductPersistenceInterface(ctrl)
	service := application.ProductService{
		ProductPersistence: mockPersistence,
	}

	product := application.NewProduct("Product 1", 10)

	t.Run("Product exists", func(t *testing.T) {
		mockPersistence.EXPECT().Get(gomock.Any()).Return(product, nil).Times(1)

		result, err := service.Get("abc")
		assert.Nil(t, err)
		assert.Equal(t, product, result)
	})

	t.Run("Product persistence throw error", func(t *testing.T) {
		mockPersistence.EXPECT().Get(gomock.Any()).Return(nil, errors.New("Internal error")).Times(1)

		result, err := service.Get("abc")
		assert.Nil(t, result)
		assert.Equal(t, "Internal error", err.Error())
	})
}

func TestProductServiceCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPersistence := mock.NewMockProductPersistenceInterface(ctrl)
	service := application.ProductService{
		ProductPersistence: mockPersistence,
	}

	t.Run("Success", func(t *testing.T) {
		product := application.NewProduct("Product 2", 10)

		mockPersistence.EXPECT().Save(gomock.Any()).Return(product, nil).Times(1)

		result, err := service.Create("Product 2", 10)
		assert.Nil(t, err)
		assert.Equal(t, result, product)
	})

	t.Run("Error - Create product with negative price", func(t *testing.T) {
		result, err := service.Create("Product with error", -10)
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, "The price must be greater than or equal to zero", err.Error())
	})

	t.Run("Error - Save persistence throws an error", func(t *testing.T) {
		mockPersistence.EXPECT().Save(gomock.Any()).Return(nil, errors.New("Internal error")).Times(1)

		result, err := service.Create("Product a", 10)
		assert.Nil(t, result)
		assert.Equal(t, "Internal error", err.Error())
	})
}

func TestProductServiceEnable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPersistence := mock.NewMockProductPersistenceInterface(ctrl)
	service := application.ProductService{
		ProductPersistence: mockPersistence,
	}

	t.Run("Success", func(t *testing.T) {
		product := application.NewProduct("Product 3", 10)

		mockPersistence.EXPECT().Save(gomock.Any()).Return(product, nil).Times(1)

		result, err := service.Enable(product)
		assert.Nil(t, err)
		assert.Equal(t, result, product)
	})

	t.Run("Error - Enable product without price", func(t *testing.T) {
		product := application.NewProduct("Product 3", 0)

		result, err := service.Enable(product)
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, "The price must be greater than zero to enable the product", err.Error())
	})

	t.Run("Error - Save persistence throws an error", func(t *testing.T) {
		product := application.NewProduct("Product 3", 10)

		mockPersistence.EXPECT().Save(gomock.Any()).Return(nil, errors.New("Internal error")).Times(1)

		result, err := service.Enable(product)
		assert.Nil(t, result)
		assert.Equal(t, "Internal error", err.Error())
	})
}

func TestProductServiceDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPersistence := mock.NewMockProductPersistenceInterface(ctrl)
	service := application.ProductService{
		ProductPersistence: mockPersistence,
	}

	t.Run("Success", func(t *testing.T) {
		product := application.NewProduct("Product 4", 10)
		product.Enable()
		product.ChangePrice(0)

		mockPersistence.EXPECT().Save(gomock.Any()).Return(product, nil).Times(1)

		result, err := service.Disable(product)
		assert.Nil(t, err)
		assert.Equal(t, result, product)
	})

	t.Run("Success - Disable already disabled product", func(t *testing.T) {
		product := application.NewProduct("Product 4", 0)

		mockPersistence.EXPECT().Save(gomock.Any()).Return(product, nil).Times(1)

		result, err := service.Disable(product)
		assert.Nil(t, err)
		assert.Equal(t, result, product)
	})

	t.Run("Error - Disable product with price greater than zero", func(t *testing.T) {
		product := application.NewProduct("Product 4", 10)

		result, err := service.Disable(product)
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, "The price must be zero to disable the product", err.Error())
	})

	t.Run("Error - Save persistence throws an error", func(t *testing.T) {
		product := application.NewProduct("Product 4", 10)
		product.Enable()
		product.ChangePrice(0)

		mockPersistence.EXPECT().Save(gomock.Any()).Return(nil, errors.New("Internal error")).Times(1)

		result, err := service.Disable(product)
		assert.Nil(t, result)
		assert.Equal(t, "Internal error", err.Error())
	})
}
