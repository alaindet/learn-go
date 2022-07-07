package main

type Filter struct{}

func (f *Filter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0, len(products))

	for i, p := range products {
		if spec.IsSatisfied(&p) {
			result = append(result, &products[i])
		}
	}

	return result
}
