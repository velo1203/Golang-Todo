package routers

import (
	"studioj/boilerplate_go/internal/controllers"
	"studioj/boilerplate_go/internal/repositories"
	"studioj/boilerplate_go/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthRouter interface {
	SetRouter(db *gorm.DB, router *gin.Engine)
}

type authRouter struct {
	db              *gorm.DB
	userRepository  repositories.UserRepository
	tokenRepository repositories.TokenRepository
	service         services.AuthService
	tokenService    services.TokenService
	controller      controllers.AuthController
	router          *gin.Engine
}

func InitAuthRouter(db *gorm.DB, router *gin.Engine) {
	r := NewAuthRouter(db, router)
	r.SetAuthRouter(db, router)
}

func NewAuthRouter(db *gorm.DB, router *gin.Engine) *authRouter {
	userRepository := repositories.NewUserRepository(db)
	tokenRepository := repositories.NewTokenRepository(db)
	service := services.NewAuthService(userRepository, tokenRepository)
	tokenService := services.NewTokenService(tokenRepository)
	controller := controllers.NewAuthController(service)

	return &authRouter{
		db:              db,
		userRepository:  userRepository,
		tokenRepository: tokenRepository,
		service:         service,
		tokenService:    tokenService,
		controller:      controller,
		router:          router,
	}
}

func (r *authRouter) SetAuthRouter(db *gorm.DB, router *gin.Engine) {
	group := router.Group("/auth")
	group.POST("login", r.controller.Login)
	group.POST("register", r.controller.Register)
}
