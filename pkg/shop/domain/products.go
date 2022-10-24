package products

import (
	"errors"
	"github.com/noel/monolith-microservice/pkg/common/price"
)

type ID string

var (
	ErrEmptyID   = errors.New("empty product id")
	ErrEmptyName = errors.New("empty product name")
)

type Product struct {
	id          ID
	name        string
	description string
	price       price.Price
}

func Newproduct(id ID, name string, description string, price price.Price) (*product, error) {
	if len(id) == 0 {
		return nil, ErrEmptyID
	}
	if len(name) == 0 {
		return nil, ErrEmptyName
	}
	return &product{
		id:          "",
		name:        "",
		description: "",
		price:       price.Price{},
	}, nil
}

func (p Product) ID() ID {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p Product) Description() string {
	return p.description
}

func (p Product) Price() price.Price {
	return p.price
}
