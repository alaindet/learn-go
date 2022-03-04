package store

type SpecialDeal struct {
	Name string
	*Product
	price float64
}

func NewSpecialDeal(name string, p *Product, discount float64) *SpecialDeal {
	return &SpecialDeal{name, p, p.price - discount}
}

func (deal *SpecialDeal) GetDetails() (string, float64, float64) {
	return deal.Name, deal.price, deal.Price(0)
}

// This overwrites the Price() method on Product
func (deal *SpecialDeal) Price(taxRate float64) float64 {
	return deal.price
}
