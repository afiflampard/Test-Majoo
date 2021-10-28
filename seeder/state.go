package seeder

import "majoo/models"

func StateUser() []models.OrderState {
	var states = []models.OrderState{
		models.OrderState{
			State: "Jual",
		},
		models.OrderState{
			State: "Beli",
		},
	}
	return states
}
