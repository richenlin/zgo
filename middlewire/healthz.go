package middlewire

// 服务器健康检查
import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Healthz 路由
type Healthz gin.IRoutes

// NewHealthz 初始化根路由
func NewHealthz(app *gin.Engine) Healthz {
	return app.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    time.Now().Unix(),
		})
	})
}
