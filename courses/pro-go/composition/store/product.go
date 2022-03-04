package store

type Describable interface {
	GetName() string
	GetCategory() string
	ItemForSale // <-- This is an embedded interface
}

type Product struct {
	Name, Category string
	price          float64
}

func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (p *Product) Price(taxRate float64) float64 {
	return p.price * (1 + taxRate)
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetCategory() string {
	return p.Category
}
