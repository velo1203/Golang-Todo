package routers

import (
	"studioj/boilerplate_go/internal/controllers"
	"studioj/boilerplate_go/internal/repositories"
	"studioj/boilerplate_go/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TokenRouter interface {
	SetRouter(db *gorm.DB, router *gin.Engine)
}

type tokenRouter struct {
	db         *gorm.DB
	repository repositories.TokenRepository
	service    services.TokenService
	controller controllers.TokenController
	router     *gin.Engine
}

func InitTokenRouter(db *gorm.DB, router *gin.Engine) {
	r := NewTokenRouter(db, router)
	r.SetTokenRouter(db, router)
}

func NewTokenRouter(db *gorm.DB, router *gin.Engine) *tokenRouter {
	repository := repositories.NewTokenRepository(db)
	service := services.NewTokenService(repository)
	controller := controllers.NewTokenController(service)

	return &tokenRouter{
		db:         db,
		repository: repository,
		service:    service,
		controller: controller,
		router:     router,
	}
}

func (r *tokenRouter) SetTokenRouter(db *gorm.DB, router *gin.Engine) {
	// group := router.Group("/token")
	// group.GET("refresh", middleware.Auth(), r.controller.Refresh)
}
