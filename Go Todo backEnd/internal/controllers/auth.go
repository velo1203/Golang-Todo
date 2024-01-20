package controllers

import (
	"net/http"

	"studioj/boilerplate_go/internal/models"
	"studioj/boilerplate_go/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(*gin.Context)
	Login(*gin.Context)
}

type authController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) AuthController {
	return &authController{
		service: service,
	}
}

func (controller *authController) Register(c *gin.Context) {
	var params models.RegisterRequest
	if c.BindJSON(&params) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	user, err := controller.service.Register(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := controller.service.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user, "token": token.Token, "refresh_token": token.RefreshToken})
}

func (controller *authController) Login(c *gin.Context) {
	var params models.LoginRequest
	if c.BindJSON(&params) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	user, err := controller.service.Login(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := controller.service.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user, "token": token.Token, "refresh_token": token.RefreshToken})
}
