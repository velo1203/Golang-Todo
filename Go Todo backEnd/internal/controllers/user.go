package controllers

import (
	"fmt"
	"net/http"

	"studioj/boilerplate_go/internal/services"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Find(*gin.Context)
}

type userController struct {
	service      services.UserService
	tokenService services.TokenService
}

func NewUserController(service services.UserService, tokenService services.TokenService) UserController {
	return &userController{
		service:      service,
		tokenService: tokenService,
	}
}

func (controller *userController) Find(c *gin.Context) {
	id := c.Param("id")
	user_id := c.GetString("sub")
	fmt.Println("id", id)
	fmt.Println("user_id", user_id)

	if user_id != id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong user"})
		return
	}

	result, err := controller.service.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
