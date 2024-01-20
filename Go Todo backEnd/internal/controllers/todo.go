package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"studioj/boilerplate_go/internal/models"
	"studioj/boilerplate_go/internal/services"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	Create(*gin.Context)
	Read(*gin.Context)
	Delete(*gin.Context)
	Put(*gin.Context)
}
type todoController struct {
	service services.TodoService
}

func NewTodoController(service services.TodoService) TodoController {
	return &todoController{service: service}
}

func (controller *todoController) Create(c *gin.Context) {
	userID, exists := c.Get("sub")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var params models.CreateRequest
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	createdTodo, err := controller.service.Create(userID.(string), &params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo": createdTodo})
}

func (controller *todoController) Read(c *gin.Context) {
	idStr := c.Query("ID")
	userID, exists := c.Get("sub")
	fmt.Println("userID : ", userID)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var id uint
	if idStr != "" {
		parsedID, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}
		id = uint(parsedID)

		// 특정 투두 항목 조회
		todo, err := controller.service.ReadByID(userID.(string), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if todo == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"todo": todo})
	} else {
		// 사용자의 모든 투두 항목 조회
		todos, err := controller.service.ReadAll(userID.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"todos": todos})
	}
}
func (controller *todoController) Delete(c *gin.Context) {
	userID, exists := c.Get("sub")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var params models.DeleteRequest
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	err := controller.service.Delete(userID.(string), params.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"delete": "success"})
}

func (controller *todoController) Put(c *gin.Context) {
	userID, exists := c.Get("sub")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var params models.UpdateRequest
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	updatedTodo, err := controller.service.Update(userID.(string), &params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo": updatedTodo})
}
