package routers

import (
	"net/http"
	"studioj/boilerplate_go/internal/database"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Init() {
	gin.SetMode(gin.ReleaseMode)
	Router = gin.Default()

	// CORS 미들웨어 설정
	Router.Use(CORSMiddleware())

	maindb := database.GetDB()

	InitAuthRouter(maindb, Router)
	InitTokenRouter(maindb, Router)
	InitUserRouter(maindb, Router)
	InitTodoRouter(maindb, Router)
}

// CORSMiddleware creates a Gin middleware for CORS.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
