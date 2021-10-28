package seeder

import (
	"log"
	"majoo/models"

	"golang.org/x/crypto/bcrypt"
)

func SeedUser() ([]models.User, []models.Role) {
	pass, err := bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	var roles = []models.Role{
		{
			Role: "Merchant",
		},
		{
			Role: "Supplier",
		},
		{
			Role: "Pelanggan",
		},
	}
	var users = []models.User{
		{
			Username: "Afif",
			FullName: "Afif Musyayyidin",
			RoleID:   1,
			Password: string(pass),
		},
		{
			Username: "Fifa",
			FullName: "Afif Musyayyidin",
			RoleID:   2,
			Password: string(pass),
		},
		{
			Username: "Fif",
			FullName: "Afif Musyay",
			RoleID:   3,
			Password: string(pass),
		},
	}
	return users, roles
}
