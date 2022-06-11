package main

import "final_project/data"

func (app *Config) getInvoice(u data.User, plan *data.Plan) (string, error) {
	// TODO: Add invoice logic here if needed
	return plan.PlanAmountFormatted, nil
}
