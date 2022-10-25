package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/noel/monolith-microservice/pkg/common/price"
	products "github.com/noel/monolith-microservice/pkg/shop/domain"
	"net/http"
)

func AddRoutes(router chi.Mux, productsReadModel productsReadmodel) {
	resource := productsResource{productsReadModel}
	router.Get("/products", resource.GetAll)
}

type productsReadmodel interface {
	AllProducts() ([]products.Product, error)
}

type productsResource struct {
	readModel productsReadmodel
}

type productView struct {
	ID          string    `json:"ID"`
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	Price       priceView `json:"Price"`
}

type priceView struct {
	Cents    uint   `json:"cents"`
	Currency string `json:"currency"`
}

func priceViewFromPrice(p price.Price) priceView {
	return priceView{p.Cents(), p.Currency()}

}

func (p productsResource) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := p.readModel.AllProducts()
	if err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}

	view := []productView{}
	for _, product := range products {
		view = append(view, productView{
			string(product.ID()),
			product.Name(),
			product.Description(),
			priceViewFromPrice(product.Price()),
		})
	}
	render.Respond(w, r, view)
}
