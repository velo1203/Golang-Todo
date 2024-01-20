package routers

import (
	"studioj/boilerplate_go/internal/controllers"
	"studioj/boilerplate_go/internal/middleware"
	"studioj/boilerplate_go/internal/repositories"
	"studioj/boilerplate_go/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodoRouter struct {
	db             *gorm.DB
	todoRepository repositories.TodoRepository
	service        services.TodoService
	controller     controllers.TodoController
	router         *gin.Engine
}

func InitTodoRouter(db *gorm.DB, router *gin.Engine) {
	r := NewTodoRouter(db, router)
	r.SetTodoRoutes(db, router)
}
func NewTodoRouter(db *gorm.DB, router *gin.Engine) *TodoRouter {
	todoRepository := repositories.NewTodoRepository(db)
	service := services.NewTodoService(todoRepository)
	controller := controllers.NewTodoController(service)
	return &TodoRouter{
		db:             db,
		todoRepository: todoRepository,
		service:        service,
		controller:     controller,
		router:         router,
	}
}

func (r *TodoRouter) SetTodoRoutes(db *gorm.DB, router *gin.Engine) {
	todoGroup := router.Group("/todos").Use(middleware.Auth())
	todoGroup.POST("/", r.controller.Create)
	todoGroup.GET("/", r.controller.Read)
	todoGroup.DELETE("/", r.controller.Delete)
	todoGroup.PUT("/", r.controller.Put)
}
