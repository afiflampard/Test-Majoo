package service

import (
	"log"
	"majoo/entities"
	"majoo/helpers"
	"majoo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginUser(db *gorm.DB, user entities.Login) (*entities.LoginResponse, error) {

	var userTemp models.User
	if err := db.Where("username = ?", user.Username).Find(&userTemp).Error; err != nil {
		return nil, err
	}
	result := bcrypt.CompareHashAndPassword([]byte(userTemp.Password), []byte(user.Password))

	if result != nil {
		return nil, result
	}
	token, err := GenerateToken(&userTemp)
	if err != nil {
		return nil, err
	}
	return &entities.LoginResponse{
		IsUser: true,
		Token:  token,
	}, nil
}

func CreateUser(db *gorm.DB, user *models.User) (*helpers.Response, *models.User) {
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	addPerson := models.User{
		Username: user.Username,
		Password: string(pass),
		RoleID:   user.RoleID,
	}
	if err := db.Create(&addPerson).Error; err != nil {
		return &helpers.Response{
			Kode:    http.StatusBadRequest,
			Message: "Cannot add User",
			Status:  false,
		}, nil
	}
	return &helpers.Response{
		Kode:    200,
		Message: "User Success created",
		Status:  true,
	}, &addPerson
}

func GetUserByID(db *gorm.DB, id int) (*models.User, error) {
	var user *models.User
	if err := db.Preload("Role").First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetAllUser(db *gorm.DB, c *gin.Context) (user []models.User, err error) {
	if err := db.Preload("Role").Find(&user).Error; err != nil {
		helpers.Responses(c, http.StatusBadRequest, "User Not Found", false)
	}
	return user, nil
}

func UpdateUser(db *gorm.DB, id int, user *models.User) (*models.User, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	var updateUser models.User
	if err := db.Where("id = ?", id).First(&updateUser).Error; err != nil {
		return nil, err
	}
	if user.Username != "" {
		updateUser.Username = user.Username
		updateUser.Password = string(pass)
		updateUser.RoleID = user.RoleID
		db.Save(&updateUser)
	}
	return &updateUser, nil
}

func DeleteUser(db *gorm.DB, id int) (map[string]string, error) {
	var user models.User
	if err := db.Delete(&user, id).Error; err != nil {
		return map[string]string{
			"message": "User tidak ada",
		}, err
	}
	return map[string]string{
		"message": "User telah terhapus",
	}, nil
}
