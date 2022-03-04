package store

type Boat struct { // <-- The enclosing type
	*Product  // <-- An embedded field
	Capacity  int
	Motorized bool
}

func NewBoat(name string, price float64, capacity int, motorized bool) *Boat {
	product := NewProduct(name, "Watersports", price)
	return &Boat{product, capacity, motorized}
}
