package migrations

import (
	"log"
	"majoo/models"
	"majoo/seeder"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	tableExist := (db.Migrator().HasTable(&models.User{}) && db.Migrator().HasTable(&models.Outlet{}) && db.Migrator().HasTable(&models.Product{}) &&
		db.Migrator().HasTable(&models.Order{}) && db.Migrator().HasTable(&models.OrderState{}) && db.Migrator().HasTable(&models.History{}) && db.Migrator().HasTable(&models.Role{}))
	if !tableExist {
		dbMigrate := db.Debug().Migrator().DropTable(&models.User{}, &models.Outlet{}, &models.Product{}, &models.Order{}, &models.OrderState{}, &models.History{}, &models.Role{})
		if dbMigrate != nil {
			log.Fatal("Cannot Drop Table")
		}
		db.AutoMigrate(&models.User{})
		users, roles := seeder.SeedUser()
		states := seeder.StateUser()
		for _, role := range roles {
			err := db.Debug().Create(&role).Error
			if err != nil {
				log.Fatalf("Failed to create role")
			}
		}

		for _, user := range users {
			err := db.Debug().Create(&user).Error
			if err != nil {
				log.Fatalf("Failed to create user")
			}
		}
		for _, state := range states {
			err := db.Debug().Create(&state).Error
			if err != nil {
				log.Fatalf("Failed to create state")
			}
		}
	}
}
