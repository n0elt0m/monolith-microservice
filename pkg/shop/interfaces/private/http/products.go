package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/noel/monolith-microservice/pkg/common/price"
	"net/http"
)

func AddRoutes(router chi.Mux, repo products_domain.Repository) {
	resource := productsResource{repo}
	router.Get("/products/{id}", resource.Get)
}

type productsResource struct {
	repo products_domain.Repository
}

type priceView struct {
	Cents    uint   `json:"cents"`
	Currency string `json:"currency"`
}

func priceViewFromPrice(p price.Price) priceView {
	return priceView{
		Cents:    p.Cents(),
		Currency: p.Currency(),
	}
}

func (p productsResource) Get(w http.ResponseWriter, r *http.Request) {
	product, err := p.repo.ByID(products_domain.ID(chi.URLParam(r, "id")))
	if err != nil {
		_ = render.Render(w, rcommon_http.ErrInternal(err))
		return
	}

	render.Respond(w, r, ProductView{
		string(product.ID()),
		product.Name(),
		product.Description(),
		priceViewFromPrice(product.Price()),
	})
}
