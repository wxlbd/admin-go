//go:build wireinject
// +build wireinject

package main

import (
	"github.com/wxlbd/admin-go/internal/api/handler"
	"github.com/wxlbd/admin-go/internal/api/router"
	"github.com/wxlbd/admin-go/internal/middleware"
	"github.com/wxlbd/admin-go/internal/pkg/permission"
	"github.com/wxlbd/admin-go/internal/pkg/websocket"
	"github.com/wxlbd/admin-go/internal/repo"
	"github.com/wxlbd/admin-go/internal/service/infra"
	"github.com/wxlbd/admin-go/internal/service/system"

	"github.com/wxlbd/admin-go/pkg/cache"
	"github.com/wxlbd/admin-go/pkg/database"
	"github.com/wxlbd/admin-go/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitApp() (*gin.Engine, error) {
	wire.Build(
		database.InitDB,
		cache.InitRedis,
		logger.NewLogger,
		// Repo (GORM Gen)
		repo.NewQuery,
		// WebSocket
		websocket.NewManager,
		// System Services
		system.NewOAuth2TokenService,
		system.NewSmsClientFactory,
		system.NewSmsChannelService,
		system.NewSmsTemplateService,
		system.NewSmsLogService,
		system.NewSmsSendService,
		system.NewSmsCodeService,
		system.NewSocialUserService,
		system.NewAuthService,
		system.NewMenuService,
		system.NewRoleService,
		system.NewPermissionService,
		system.NewTenantService,
		system.NewTenantPackageService,
		system.NewUserService,
		system.NewDictService,
		system.NewDeptService,
		system.NewPostService,
		system.NewNoticeService,
		system.NewNotifyService,
		system.NewMailService,
		system.NewConfigService,
		system.NewLoginLogService,
		system.NewOperateLogService,
		// Infra Services
		infra.NewFileConfigService,
		infra.NewFileService,
		infra.NewScheduler,
		infra.NewJobService,
		infra.NewJobLogService,
		infra.NewApiAccessLogService,
		infra.NewApiErrorLogService,
		// Handlers
		handler.ProviderSet,
		// Casbin & Middleware
		permission.InitEnforcer,
		middleware.NewCasbinMiddleware,
		// Router
		router.InitRouter,
		// Job Handlers (Empty for basic admin template)
		ProvideJobHandlers,
	)
	return &gin.Engine{}, nil
}

// ProvideJobHandlers 聚合定时任务处理器，当前为空（管理后台基础模板）
func ProvideJobHandlers() []infra.JobHandler {
	return []infra.JobHandler{}
}
