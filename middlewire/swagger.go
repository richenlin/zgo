package middlewire

// 服务器swagger
/*
Package app 生成swagger文档

文档规则请参考：https://github.com/swaggo/swag#declarative-comments-format

使用方式：

	go get -u github.com/swaggo/swag/cmd/swag
	make swagger

*/
import (
	"zgo/modules/config"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

// Swagger swagger
type Swagger gin.IRoutes

// NewSwagger swagger
func NewSwagger(app *gin.Engine) Swagger {
	if config.C.Swagger {
		return app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return nil
}
