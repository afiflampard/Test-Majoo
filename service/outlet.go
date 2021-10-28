package service

import (
	"errors"
	"majoo/models"

	"gorm.io/gorm"
)

func CreateOutlet(db *gorm.DB, idUser int) (*models.Outlet, error) {
	var user models.User
	if err := db.Debug().First(&user, idUser).Preload("User").Preload("User.Role").Error; err != nil {
		return nil, err
	}
	if user.RoleID == 1 {
		newOutlet := models.Outlet{
			UserId: uint(idUser),
			Photo:  user.Photo,
		}
		if err := db.Create(&newOutlet).Error; err != nil {
			return nil, err
		}
		if newOutlet.ID == 1 {
			newOutlet.ID = 2
			if err := db.Save(&newOutlet).Error; err != nil {
				return nil, err
			}
		}
		return &newOutlet, nil
	}
	return nil, errors.New("You cannot add Outlet")
}
