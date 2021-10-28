package controllers

import (
	"majoo/entities"
	"majoo/models"
	"majoo/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct{}

func GetID(c *gin.Context) float64 {
	return c.MustGet("user_id").(float64)
}

func UserControllers() UserController {
	return User{}
}

func (ctx User) Login(c *gin.Context) {
	userLogin := &entities.Login{}
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Request tidak valid")
		return
	}
	resp, err := service.LoginUser(GetDB(), *userLogin)
	if err != nil {
		c.JSON(http.StatusBadGateway, map[string]string{
			"message": "Cannot Login Username or Password wrong",
		})
		return
	}
	c.JSON(http.StatusAccepted, resp)

}
func (ctx User) Delete(c *gin.Context) {
	id := c.Param("id")
	idUser, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Masukkan id")
		return
	}
	resp, err := service.DeleteUser(GetDB(), idUser)
	if err != nil {
		c.JSON(http.StatusBadGateway, resp)
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

func (ctx User) FindAll(c *gin.Context) {
	resp, err := service.GetAllUser(GetDB(), c)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "User Not Found",
		})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}
func (ctx User) FindById(c *gin.Context) {
	id := c.Param("id")
	idUser, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Masukkan id")
		return
	}
	resp, err := service.GetUserByID(GetDB(), idUser)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "User Not Found",
		})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}
func (ctx User) SignUp(c *gin.Context) {
	newUser := models.User{}
	if err := c.ShouldBindJSON(&newUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, "User bad request")
	}
	resp, err := service.CreateUser(GetDB(), &newUser)
	if err != nil {
		c.JSON(http.StatusBadGateway, resp)
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

func (ctx User) Update(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	updateUser := models.User{}
	if err != nil {
		c.JSON(http.StatusBadRequest, "Request not valid")
		return
	}
	if err := c.ShouldBindJSON(&updateUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, "Request not valid")
		return
	}
	resp, err := service.UpdateUser(GetDB(), userId, &updateUser)
	if err != nil {
		c.JSON(http.StatusBadGateway, map[string]string{
			"message": "user cannot update",
		})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}
