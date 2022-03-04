package store

type Crew struct {
	Captain, FirstOfficer string
}

type RentalBoat struct {
	*Boat
	IncludeCrew bool
	*Crew
}

func NewRentalBoat(
	name string,
	price float64,
	capacity int,
	motorized,
	crewed bool,
	captain, firstOfficer string,
) *RentalBoat {
	boat := NewBoat(name, price, capacity, motorized)
	crew := &Crew{captain, firstOfficer}
	return &RentalBoat{boat, crewed, crew}
}
