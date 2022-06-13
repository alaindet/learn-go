package data

import (
	"fmt"
	"time"
)

type TestPlan struct {
	ID                  int
	PlanName            string
	PlanAmount          int
	PlanAmountFormatted string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func testGetDummyPlan() *Plan {
	return &Plan{
		ID:         1,
		PlanName:   "Bronze Plan",
		PlanAmount: 1000,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (p *TestPlan) GetAll() ([]*Plan, error) {
	plans := []*Plan{
		testGetDummyPlan(),
	}

	return plans, nil
}

func (p *TestPlan) GetOne(id int) (*Plan, error) {
	return testGetDummyPlan(), nil
}

func (p *TestPlan) SubscribeUserToPlan(user User, plan Plan) error {
	return nil
}

// AmountForDisplay formats the price we have in the DB as a currency string
func (p *TestPlan) AmountForDisplay() string {
	amount := float64(p.PlanAmount) / 100.0
	return fmt.Sprintf("$%.2f", amount)
}
