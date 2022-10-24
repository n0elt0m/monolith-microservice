package products

import "errors"

var ErrnotFound = errors.New("product not found")

type Repository interface {
	save(*product) error
	ByID(ID) (*product, error)
}
