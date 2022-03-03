package main

type Service struct {
	description    string
	durationMonths int
	monthlyFee     float64
	features       []string
}

// Implementing Expense interface
func (s Service) getName() string {
	return s.description
}

// Implementing Expense interface
func (s Service) getCost(recur bool) float64 {
	if recur {
		return s.monthlyFee * float64(s.durationMonths)
	}
	return s.monthlyFee
}
