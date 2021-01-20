package seeder

import (
	"helloworld/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func SeedUser() []models.User {
	pass, err := bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	var users = []models.User{
		models.User{
			Username: "sayidin",
			Password: string(pass),
			FullName: "Afif Musyayyidin",
		},
		models.User{
			Username: "Fifa",
			Password: string(pass),
			FullName: "Fifa Musyayyidin",
		},
	}
	return users
}
