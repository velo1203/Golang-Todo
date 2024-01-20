package routers

import (
	"studioj/boilerplate_go/internal/controllers"
	"studioj/boilerplate_go/internal/middleware"
	"studioj/boilerplate_go/internal/repositories"
	"studioj/boilerplate_go/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRouter interface {
	SetRouter(db *gorm.DB, router *gin.Engine)
}

type userRouter struct {
	db         *gorm.DB
	repository repositories.UserRepository
	service    services.UserService
	controller controllers.UserController
	router     *gin.Engine
}

func InitUserRouter(db *gorm.DB, router *gin.Engine) {
	r := NewUserRouter(db, router)
	r.SetUserRouter(db, router)
}

func NewUserRouter(db *gorm.DB, router *gin.Engine) *userRouter {
	repository := repositories.NewUserRepository(db)
	tokenRepository := repositories.NewTokenRepository(db)
	service := services.NewUserService(repository, tokenRepository)

	tokenService := services.NewTokenService(tokenRepository)
	controller := controllers.NewUserController(service, tokenService)

	return &userRouter{
		db:         db,
		repository: repository,
		service:    service,
		controller: controller,
		router:     router,
	}
}

func (r *userRouter) SetUserRouter(db *gorm.DB, router *gin.Engine) {
	group := router.Group("/user")
	group.GET("/:id", middleware.Auth(), r.controller.Find)
}
