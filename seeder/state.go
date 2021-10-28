package seeder

import "majoo/models"

func StateUser() []models.OrderState {
	var states = []models.OrderState{
		{
			State: "Jual",
		},
		{
			State: "Beli",
		},
	}
	return states
}
