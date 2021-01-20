package migrations

import (
	"helloworld/models"
	"helloworld/seeder"
	"log"

	"gorm.io/gorm"
)

var users = seeder.SeedUser()

func User(db *gorm.DB) error {
	dbUser := db.Debug().Migrator().DropTable(&models.User{})
	if dbUser != nil {
		log.Fatal("Cannot Drop table")
	}
	dbUser = db.Debug().AutoMigrate(&models.User{})
	if dbUser != nil {
		log.Fatal("Cannot migrate Table")
	}

	for _, user := range users {
		err := db.Debug().Create(&user).Error
		if err != nil {
			log.Fatalf("failde to create user")
		}
	}

	return nil
}
