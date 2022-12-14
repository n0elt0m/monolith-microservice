package orders

import (
	"errors"
	"github.com/noel/monolith-microservice/pkg/common/price"
)

//import "github.com/noel/monolith-microservice/pkg/common/price"

type ProductID string

var ErrEmptyProductID = errors.New("empty product id")

type Product struct {
	id    ProductID
	name  string
	price price.Price
}

func NewProduct(id ProductID, name string, price price.Price) (Product, error) {
	if len(id) == 0 {
		return Product{}, ErrEmptyProductID
	}
	return Product{
		id:    id,
		name:  name,
		price: price,
	}, nil
}

func (p Product) ID() ProductID {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p Product) Price() price.Price {
	return p.price
}
