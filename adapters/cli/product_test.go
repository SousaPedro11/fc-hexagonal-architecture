package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sousapedro11/fc-arquitetura-hexagonal/adapters/cli"
	"github.com/sousapedro11/fc-arquitetura-hexagonal/application/mock"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product 1"
	productPrice := 19.99
	productStatus := "enabled"
	productId := "681051e4-2936-4b4c-87a4-efaf7b8c02ba"

	productMock := mock.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	serviceMock := mock.NewMockProductServiceInterface(ctrl)
	serviceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Enable(productMock).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Disable(productMock).Return(productMock, nil).AnyTimes()

	tests := []struct {
		price    float64
		expected string
		testName string
		name     string
		id       string
		status   string
		action   string
		err      bool
	}{
		{
			testName: "Success - Create",
			price:    productPrice,
			name:     productName,
			id:       "",
			status:   "",
			action:   "create",
			err:      false,
			expected: fmt.Sprintf("Product Id %s with the name %s has been created with the price %f and status %s", productId, productName, productPrice, productStatus),
		},
		{
			testName: "Success - Enable",
			price:    0,
			name:     "",
			id:       productId,
			status:   "",
			action:   "enable",
			err:      false,
			expected: fmt.Sprintf("Product %s has been enabled", productName),
		},
		{
			testName: "Success - Disable",
			price:    0,
			name:     "",
			id:       productId,
			status:   "",
			action:   "disable",
			err:      false,
			expected: fmt.Sprintf("Product %s has been enabled", productName),
		},
		{
			testName: "Success - Get",
			price:    0,
			name:     "",
			id:       productId,
			status:   "",
			action:   "get",
			err:      false,
			expected: fmt.Sprintf("Product Id: %s\nName: %s\nPrice: %f\nStatus: %s", productId, productName, productPrice, productStatus),
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			result, err := cli.Run(serviceMock, tt.action, tt.id, tt.name, tt.price)
			assert.Equal(t, tt.err, err != nil)
			assert.Equal(t, tt.expected, result)
		})
	}
}
