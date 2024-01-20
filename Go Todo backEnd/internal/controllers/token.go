package controllers

import (
	"net/http"

	"studioj/boilerplate_go/internal/services"

	"github.com/gin-gonic/gin"
)

type TokenController interface {
	Find(*gin.Context)
}

type tokenController struct {
	service services.TokenService
}

func NewTokenController(service services.TokenService) TokenController {
	return &tokenController{
		service: service,
	}
}

func (controller *tokenController) Find(c *gin.Context) {
	id := c.Param("id")
	user_id := c.GetString("user_id")

	if user_id != id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong user"})
		return
	}

	result, err := controller.service.Find(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (controller *tokenController) List(c *gin.Context) {
	id := c.Param("id")
	user_id := c.GetString("user_id")

	if user_id != id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong user"})
		return
	}

	result, err := controller.service.Find(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (controller *tokenController) Refresh(c *gin.Context) {
}