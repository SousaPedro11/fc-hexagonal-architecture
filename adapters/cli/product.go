package cli

import (
	"fmt"

	"github.com/sousapedro11/fc-arquitetura-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action, productId, producName string, productPrice float64) (string, error) {
	result := ""

	switch action {
	case "create":
		product, err := service.Create(producName, productPrice)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product Id %s with the name %s has been created with the price %f and status %s", product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		product, err = service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been enabled", product.GetName())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		product, err = service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been enabled", product.GetName())
	default:
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product Id: %s\nName: %s\nPrice: %f\nStatus: %s", product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	}

	return result, nil
}
