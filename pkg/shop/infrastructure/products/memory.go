package products

import products "github.com/noel/monolith-microservice/pkg/shop/domain"

type MemoryRepository struct {
	products []products.Product
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{[]products.Product{}}
}

func (m *MemoryRepository) Save(productTosave *products.Product) error {
	for i, p := range m.products {
		if p.ID() == productTosave.ID() {
			m.products[i] = *productTosave
			return nil
		}
	}
	m.products = append(m.products, *productTosave)
	return nil
}

func (m MemoryRepository) ByID(id products.ID) (*products.Product, error) {
	for _, p := range m.products {
		if p.ID() == id {
			return &p, nil
		}
	}
	return nil, products.ErrnotFound
}

func (m MemoryRepository) AllProducts() ([]products.Product, error) {
	return m.products, nil
}
