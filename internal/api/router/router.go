package router

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/wxlbd/admin-go/internal/api/handler/admin"
	"github.com/wxlbd/admin-go/internal/middleware"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, rdb *redis.Client,
	adminHandlers *admin.AdminHandlers,
	casbinMiddleware *middleware.CasbinMiddleware,
) *gin.Engine {
	// Debug log to confirm router init
	fmt.Println("Initializing Router...")
	r := gin.New()
	r.Use(middleware.Recovery())
	r.Use(middleware.ErrorHandler())
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"*"},
		AllowHeaders:    []string{"*"},
	}))
	r.Use(gin.Logger())
	// 注入 gin.Context 到 request context，供 GORM Hook 使用
	r.Use(middleware.InjectContext())

	// 基础路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// ========== 模块化路由注册 ==========

	// System 模块
	// WebSocket (Register at root /infra/ws to match Java path)
	r.GET("/infra/ws", adminHandlers.Infra.WebSocket.Handle) // Corrected WebSocketHandler reference

	// System 模块
	RegisterSystemRoutes(r, adminHandlers.System, adminHandlers.Infra, casbinMiddleware)

	// Area 地区路由
	RegisterAreaRoutes(r, adminHandlers.System.Area)

	return r
}
