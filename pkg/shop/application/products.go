package application

import (
	"errors"
	"github.com/noel/monolith-microservice/pkg/common/price"
	products "github.com/noel/monolith-microservice/pkg/shop/domain"
)

type productReadmodel interface {
	AllProducts() ([]products.Product, err)
}

type ProductService struct {
	repo      products.Repository
	readModel productReadmodel
}

func NewProductsService(repo products.Repository, readmodel productReadmodel) ProductService {
return ProductService{
	repo:      repo,
	readModel: readmodel,
}
}

func (s ProductService) AllProducts()([]products.Product,error) {
return s.readModel.AllProducts()
}

type AddproductCommand struct {
	ID            string
	Name          string
	Description   string
	PriceCents    uint
	PriceCurrency string
}

func (s ProductService) AddProduct(cmd AddproductCommand) error {
	price,err:= price.NewPrice(cmd.PriceCents, cmd.PriceCurrency)
	if err != nil {
		return errors.Wrap(err,"invalid product price")
	}

	p,err:= products.NewProduct(products.ID(cmd.ID)), cmd.Name, cmd.Description, cmd.PriceCents)
	if err != nil {
		return errors.Wrap(err,"cannot create product")
	}

if err:= s.repo.Save(p);err!=nil{
	return errors.Wrap(err,"cannot save product")
}
return nil
}
