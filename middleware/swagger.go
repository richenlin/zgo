package middleware

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// SwaggerRoutes swagger请求路由
func SwaggerRoutes() gin.HandlerFunc {
	return ginSwagger.WrapHandler(swaggerFiles.Handler)
}
