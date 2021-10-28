package service

import (
	"errors"
	"majoo/models"

	"gorm.io/gorm"
)

func CreateOutlet(db *gorm.DB, idUser int) (*models.Outlet, error) {
	var user models.User
	if err := db.First(user, idUser).Preload("Role").Error; err != nil {
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
		return &newOutlet, nil
	}
	return nil, errors.New("You cannot add Outlet")
}
