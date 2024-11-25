package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
	ChangePrice(price float64) error
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReaderInterface interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriterInterface interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductWriterInterface
	ProductReaderInterface
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	Price  float64 `valid:"float,optional"`
	Id     string  `valid:"uuid"`
	Name   string  `valid:"required"`
	Status string  `valid:"required,in(disabled|enabled)"`
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		Id:     uuid.NewString(),
		Name:   name,
		Status: DISABLED,
		Price:  price,
	}
}

func (p *Product) IsValid() (bool, error) {
	if p.Price < 0 {
		return false, errors.New("The price must be greater than or equal to zero")
	}
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Product) Enable() error {
	if p.Price <= 0 {
		return errors.New("The price must be greater than zero to enable the product")
	}

	if p.Status == DISABLED {
		p.Status = ENABLED
	}
	return nil
}

func (p *Product) Disable() error {
	if p.Price > 0 {
		return errors.New("The price must be zero to disable the product")
	}
	if p.Status == ENABLED {
		p.Status = DISABLED
	}
	return nil
}

func (p *Product) GetId() string {
	return p.Id
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) ChangePrice(price float64) error {
	if price < 0 {
		return errors.New("The price must be greater than or equal to zero")
	}
	p.Price = price
	return nil
}
