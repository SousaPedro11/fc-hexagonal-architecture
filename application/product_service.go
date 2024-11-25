package application

type ProductService struct {
	ProductPersistence ProductPersistenceInterface
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.ProductPersistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct(name, price)
	if valid, err := product.IsValid(); !valid {
		return nil, err
	}
	result, err := s.ProductPersistence.Save(product)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	if err := product.Enable(); err != nil {
		return nil, err
	}
	result, err := s.ProductPersistence.Save(product)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	if err := product.Disable(); err != nil {
		return nil, err
	}
	result, err := s.ProductPersistence.Save(product)
	if err != nil {
		return nil, err
	}
	return result, nil
}
