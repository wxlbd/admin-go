package router

import (
	"github.com/wxlbd/admin-go/internal/api/handler/admin/infra"
	"github.com/wxlbd/admin-go/internal/api/handler/admin/system"
	"github.com/wxlbd/admin-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterSystemRoutes(engine *gin.Engine,
	handlers *system.Handlers,
	infraHandlers *infra.Handlers,
	casbinMiddleware *middleware.CasbinMiddleware,
) {
	api := engine.Group("/admin-api")
	{
		systemGroup := api.Group("/system")
		{
			// ====== Public Routes (No Auth Required) ======
			// Auth Public Routes
			authGroup := systemGroup.Group("/auth")
			{
				authGroup.POST("/login", handlers.Auth.Login)
				authGroup.POST("/logout", handlers.Auth.Logout)
				authGroup.POST("/refresh-token", handlers.Auth.RefreshToken)
				authGroup.POST("/sms-login", handlers.Auth.SmsLogin)
				authGroup.POST("/send-sms-code", handlers.Auth.SendSmsCode)
				authGroup.POST("/register", handlers.Auth.Register)
				authGroup.POST("/reset-password", handlers.Auth.ResetPassword)
				authGroup.GET("/social-auth-redirect", handlers.Auth.SocialAuthRedirect)
				authGroup.POST("/social-login", handlers.Auth.SocialLogin)
			}

			// Tenant Public Routes
			tenantPublicGroup := systemGroup.Group("/tenant")
			{
				tenantPublicGroup.GET("/simple-list", handlers.Tenant.GetTenantSimpleList)
				tenantPublicGroup.GET("/get-by-website", handlers.Tenant.GetTenantByWebsite)
				tenantPublicGroup.GET("/get-id-by-name", handlers.Tenant.GetTenantIdByName)
			}

			// Dict Public Routes
			dictTypePublicGroup := systemGroup.Group("/dict-type")
			{
				dictTypePublicGroup.GET("/simple-list", handlers.Dict.GetSimpleDictTypeList)
				dictTypePublicGroup.GET("/list-all-simple", handlers.Dict.GetSimpleDictTypeList)
			}

			dictDataPublicGroup := systemGroup.Group("/dict-data")
			{
				dictDataPublicGroup.GET("/simple-list", handlers.Dict.GetSimpleDictDataList)
				dictDataPublicGroup.GET("/list-all-simple", handlers.Dict.GetSimpleDictDataList)
			}

			// Dept Public Routes
			deptPublicGroup := systemGroup.Group("/dept")
			{
				deptPublicGroup.GET("/list", handlers.Dept.GetDeptList)
				deptPublicGroup.GET("/list-all-simple", handlers.Dept.GetSimpleDeptList)
				deptPublicGroup.GET("/simple-list", handlers.Dept.GetSimpleDeptList)
			}

			// Post Public Routes
			postPublicGroup := systemGroup.Group("/post")
			{
				postPublicGroup.GET("/simple-list", handlers.Post.GetSimplePostList)
			}

			// User Public Routes
			userPublicGroup := systemGroup.Group("/user")
			{
				userPublicGroup.GET("/list-all-simple", handlers.User.GetSimpleUserList)
				userPublicGroup.GET("/simple-list", handlers.User.GetSimpleUserList)
			}

			// Role Public Routes
			rolePublicGroup := systemGroup.Group("/role")
			{
				rolePublicGroup.GET("/list-all-simple", handlers.Role.GetSimpleRoleList)
				rolePublicGroup.GET("/simple-list", handlers.Role.GetSimpleRoleList)
			}

			// Menu Public Routes
			menuPublicGroup := systemGroup.Group("/menu")
			{
				menuPublicGroup.GET("/simple-list", handlers.Menu.GetSimpleMenuList)
			}

			// ====== Protected Routes (Auth Required) ======
			// Apply Auth Middleware to all subsequent system routes
			systemGroup.Use(middleware.Auth())

			// Auth Protected Routes
			authProtectedGroup := systemGroup.Group("/auth")
			{
				authProtectedGroup.GET("/get-permission-info", handlers.Auth.GetPermissionInfo)
			}

			// Tenant Protected Routes
			tenantProtectedGroup := systemGroup.Group("/tenant")
			{
				tenantProtectedGroup.POST("/create", casbinMiddleware.RequirePermission("system:tenant:create"), handlers.Tenant.CreateTenant)
				tenantProtectedGroup.PUT("/update", casbinMiddleware.RequirePermission("system:tenant:update"), handlers.Tenant.UpdateTenant)
				tenantProtectedGroup.DELETE("/delete", casbinMiddleware.RequirePermission("system:tenant:delete"), handlers.Tenant.DeleteTenant)
				tenantProtectedGroup.DELETE("/delete-list", casbinMiddleware.RequirePermission("system:tenant:delete"), handlers.Tenant.DeleteTenantList)
				tenantProtectedGroup.GET("/get", casbinMiddleware.RequirePermission("system:tenant:query"), handlers.Tenant.GetTenant)
				tenantProtectedGroup.GET("/page", casbinMiddleware.RequirePermission("system:tenant:query"), handlers.Tenant.GetTenantPage)
				tenantProtectedGroup.GET("/export-excel", casbinMiddleware.RequirePermission("system:tenant:export"), handlers.Tenant.ExportTenantExcel)
			}

			// Tenant Package Protected Routes
			tenantPackageGroup := systemGroup.Group("/tenant-package")
			{
				tenantPackageGroup.POST("/create", casbinMiddleware.RequirePermission("system:tenant-package:create"), handlers.TenantPackage.CreateTenantPackage)
				tenantPackageGroup.PUT("/update", casbinMiddleware.RequirePermission("system:tenant-package:update"), handlers.TenantPackage.UpdateTenantPackage)
				tenantPackageGroup.DELETE("/delete", casbinMiddleware.RequirePermission("system:tenant-package:delete"), handlers.TenantPackage.DeleteTenantPackage)
				tenantPackageGroup.DELETE("/delete-list", casbinMiddleware.RequirePermission("system:tenant-package:delete"), handlers.TenantPackage.DeleteTenantPackageList)
				tenantPackageGroup.GET("/get", casbinMiddleware.RequirePermission("system:tenant-package:query"), handlers.TenantPackage.GetTenantPackage)
				tenantPackageGroup.GET("/page", casbinMiddleware.RequirePermission("system:tenant-package:query"), handlers.TenantPackage.GetTenantPackagePage)
				tenantPackageGroup.GET("/get-simple-list", handlers.TenantPackage.GetTenantPackageSimpleList)
				tenantPackageGroup.GET("/simple-list", handlers.TenantPackage.GetTenantPackageSimpleList)
			}

			// Dict Type Protected Routes
			dictTypeProtectedGroup := systemGroup.Group("/dict-type")
			{
				dictTypeProtectedGroup.GET("/page", casbinMiddleware.RequirePermission("system:dict:query"), handlers.Dict.GetDictTypePage)
				dictTypeProtectedGroup.GET("/get", casbinMiddleware.RequirePermission("system:dict:query"), handlers.Dict.GetDictType)
				dictTypeProtectedGroup.POST("/create", casbinMiddleware.RequirePermission("system:dict:create"), handlers.Dict.CreateDictType)
				dictTypeProtectedGroup.PUT("/update", casbinMiddleware.RequirePermission("system:dict:update"), handlers.Dict.UpdateDictType)
				dictTypeProtectedGroup.DELETE("/delete", casbinMiddleware.RequirePermission("system:dict:delete"), handlers.Dict.DeleteDictType)
				dictTypeProtectedGroup.GET("/export-excel", casbinMiddleware.RequirePermission("system:dict:export"), handlers.Dict.ExportDictTypeExcel)
			}

			// Dict Data Protected Routes
			dictDataProtectedGroup := systemGroup.Group("/dict-data")
			{
				dictDataProtectedGroup.GET("/page", casbinMiddleware.RequirePermission("system:dict:query"), handlers.Dict.GetDictDataPage)
				dictDataProtectedGroup.GET("/get", casbinMiddleware.RequirePermission("system:dict:query"), handlers.Dict.GetDictData)
				dictDataProtectedGroup.POST("/create", casbinMiddleware.RequirePermission("system:dict:create"), handlers.Dict.CreateDictData)
				dictDataProtectedGroup.PUT("/update", casbinMiddleware.RequirePermission("system:dict:update"), handlers.Dict.UpdateDictData)
				dictDataProtectedGroup.DELETE("/delete", casbinMiddleware.RequirePermission("system:dict:delete"), handlers.Dict.DeleteDictData)
			}

			// Dept Protected Routes
			deptProtectedGroup := systemGroup.Group("/dept")
			{
				deptProtectedGroup.GET("/get", casbinMiddleware.RequirePermission("system:dept:query"), handlers.Dept.GetDept)
				deptProtectedGroup.POST("/create", casbinMiddleware.RequirePermission("system:dept:create"), handlers.Dept.CreateDept)
				deptProtectedGroup.PUT("/update", casbinMiddleware.RequirePermission("system:dept:update"), handlers.Dept.UpdateDept)
				deptProtectedGroup.DELETE("/delete", casbinMiddleware.RequirePermission("system:dept:delete"), handlers.Dept.DeleteDept)
			}

			// Post Protected Routes
			postProtectedGroup := systemGroup.Group("/post")
			{
				postProtectedGroup.GET("/page", casbinMiddleware.RequirePermission("system:post:query"), handlers.Post.GetPostPage)
				postProtectedGroup.GET("/get", casbinMiddleware.RequirePermission("system:post:query"), handlers.Post.GetPost)
				postProtectedGroup.POST("/create", casbinMiddleware.RequirePermission("system:post:create"), handlers.Post.CreatePost)
				postProtectedGroup.PUT("/update", casbinMiddleware.RequirePermission("system:post:update"), handlers.Post.UpdatePost)
				postProtectedGroup.DELETE("/delete", casbinMiddleware.RequirePermission("system:post:delete"), handlers.Post.DeletePost)
			}

			// User Protected Routes
			userGroup := systemGroup.Group("/user")
			{
				userGroup.GET("/page", casbinMiddleware.RequirePermission("system:user:query"), handlers.User.GetUserPage)
				userGroup.GET("/get", casbinMiddleware.RequirePermission("system:user:query"), handlers.User.GetUser)
				userGroup.POST("/create", casbinMiddleware.RequirePermission("system:user:create"), handlers.User.CreateUser)
				userGroup.PUT("/update", casbinMiddleware.RequirePermission("system:user:update"), handlers.User.UpdateUser)
				userGroup.DELETE("/delete", casbinMiddleware.RequirePermission("system:user:delete"), handlers.User.DeleteUser)
				userGroup.DELETE("/delete-list", casbinMiddleware.RequirePermission("system:user:delete"), handlers.User.DeleteUserList)
				userGroup.PUT("/update-status", casbinMiddleware.RequirePermission("system:user:update"), handlers.User.UpdateUserStatus)
				userGroup.PUT("/update-password", casbinMiddleware.RequirePermission("system:user:update-password"), handlers.User.UpdateUserPassword)
				userGroup.GET("/export", casbinMiddleware.RequirePermission("system:user:export"), handlers.User.ExportUser)
				userGroup.GET("/get-import-template", casbinMiddleware.RequirePermission("system:user:import"), handlers.User.GetImportTemplate)
				userGroup.POST("/import", casbinMiddleware.RequirePermission("system:user:import"), handlers.User.ImportUser)

				userProtectedGroup := userGroup.Group("")
				userProtectedGroup.Use(middleware.Auth())
				{
					// userProtectedGroup.PUT("/profile/update", handlers.User.UpdateUserProfile)
					// userProtectedGroup.PUT("/profile/update-password", handlers.User.UpdateUserProfilePassword)
					// userProtectedGroup.GET("/profile/get", handlers.User.GetUserProfile)
				}
			}

			// Role Protected Routes
			roleProtectedGroup := systemGroup.Group("/role")
			{
				roleProtectedGroup.GET("/page", casbinMiddleware.RequirePermission("system:role:query"), handlers.Role.GetRolePage)
				roleProtectedGroup.GET("/get", casbinMiddleware.RequirePermission("system:role:query"), handlers.Role.GetRole)
				roleProtectedGroup.POST("/create", casbinMiddleware.RequirePermission("system:role:create"), handlers.Role.CreateRole)
				roleProtectedGroup.PUT("/update", casbinMiddleware.RequirePermission("system:role:update"), handlers.Role.UpdateRole)
				roleProtectedGroup.PUT("/update-status", casbinMiddleware.RequirePermission("system:role:update"), handlers.Role.UpdateRoleStatus)
				roleProtectedGroup.DELETE("/delete", casbinMiddleware.RequirePermission("system:role:delete"), handlers.Role.DeleteRole)
			}

			// Permission Protected Routes
			permissionProtectedGroup := systemGroup.Group("/permission")
			{
				permissionProtectedGroup.GET("/list-role-menus", casbinMiddleware.RequirePermission("system:permission:assign-role-menu"), handlers.Permission.GetRoleMenuList)
				permissionProtectedGroup.POST("/assign-role-menu", casbinMiddleware.RequirePermission("system:permission:assign-role-menu"), handlers.Permission.AssignRoleMenu)
				permissionProtectedGroup.POST("/assign-role-data-scope", casbinMiddleware.RequirePermission("system:permission:assign-role-data-scope"), handlers.Permission.AssignRoleDataScope)
				permissionProtectedGroup.GET("/list-user-roles", casbinMiddleware.RequirePermission("system:permission:assign-user-role"), handlers.Permission.GetUserRoleList)
				permissionProtectedGroup.POST("/assign-user-role", casbinMiddleware.RequirePermission("system:permission:assign-user-role"), handlers.Permission.AssignUserRole)
			}

			// Menu Protected Routes
			menuProtectedGroup := systemGroup.Group("/menu")
			{
				menuProtectedGroup.POST("/create", casbinMiddleware.RequirePermission("system:menu:create"), handlers.Menu.CreateMenu)
				menuProtectedGroup.PUT("/update", casbinMiddleware.RequirePermission("system:menu:update"), handlers.Menu.UpdateMenu)
				menuProtectedGroup.DELETE("/delete", casbinMiddleware.RequirePermission("system:menu:delete"), handlers.Menu.DeleteMenu)
				menuProtectedGroup.GET("/list", casbinMiddleware.RequirePermission("system:menu:query"), handlers.Menu.GetMenuList)
				menuProtectedGroup.GET("/get", casbinMiddleware.RequirePermission("system:menu:query"), handlers.Menu.GetMenu)
			}

			// Notice Protected Routes
			noticeProtectedGroup := systemGroup.Group("/notice")
			{
				noticeProtectedGroup.GET("/page", casbinMiddleware.RequirePermission("system:notice:query"), handlers.Notice.GetNoticePage)
				noticeProtectedGroup.GET("/get", casbinMiddleware.RequirePermission("system:notice:query"), handlers.Notice.GetNotice)
				noticeProtectedGroup.POST("/create", casbinMiddleware.RequirePermission("system:notice:create"), handlers.Notice.CreateNotice)
				noticeProtectedGroup.PUT("/update", casbinMiddleware.RequirePermission("system:notice:update"), handlers.Notice.UpdateNotice)
				// noticeProtectedGroup.PUT("/update-status", casbinMiddleware.RequirePermission("system:notice:update"), handlers.Notice.UpdateNoticeStatus) // Commented out
				noticeProtectedGroup.DELETE("/delete", casbinMiddleware.RequirePermission("system:notice:delete"), handlers.Notice.DeleteNotice)
				noticeProtectedGroup.POST("/push", casbinMiddleware.RequirePermission("system:notice:create"), handlers.Notice.Push)
			}

			// Notify Message Protected Routes (站内信/通知消息)
			notifyMessageGroup := systemGroup.Group("/notify-message")
			{
				notifyMessageGroup.GET("/page", casbinMiddleware.RequirePermission("system:notify-message:query"), handlers.Notify.GetNotifyMessagePage)
				notifyMessageGroup.GET("/get", casbinMiddleware.RequirePermission("system:notify-message:query"), handlers.Notify.GetNotifyMessage)
				// 以下API仅需认证，不需要特定权限（对齐Java实现）- P1修复
				notifyMessageGroup.GET("/get-unread-list", middleware.Auth(), handlers.Notify.GetUnreadNotifyMessageList)
				notifyMessageGroup.GET("/get-unread-count", middleware.Auth(), handlers.Notify.GetUnreadNotifyMessageCount)
				notifyMessageGroup.PUT("/update-read", middleware.Auth(), handlers.Notify.UpdateNotifyMessageRead)
				notifyMessageGroup.PUT("/update-all-read", middleware.Auth(), handlers.Notify.UpdateAllNotifyMessageRead)
			}

			// Notify Template Protected Routes (站内信模板管理) - P2 & P3修复
			notifyTemplateGroup := systemGroup.Group("/notify-template")
			{
				notifyTemplateGroup.POST("/create",
					casbinMiddleware.RequirePermission("system:notify-template:create"),
					handlers.Notify.CreateNotifyTemplate)
				notifyTemplateGroup.PUT("/update",
					casbinMiddleware.RequirePermission("system:notify-template:update"),
					handlers.Notify.UpdateNotifyTemplate)
				notifyTemplateGroup.DELETE("/delete",
					casbinMiddleware.RequirePermission("system:notify-template:delete"),
					handlers.Notify.DeleteNotifyTemplate)
				notifyTemplateGroup.GET("/get",
					casbinMiddleware.RequirePermission("system:notify-template:query"),
					handlers.Notify.GetNotifyTemplate)
				notifyTemplateGroup.GET("/page",
					casbinMiddleware.RequirePermission("system:notify-template:query"),
					handlers.Notify.GetNotifyTemplatePage)
				// P3: SendNotify API - 对应Java NotifyTemplateController.sendNotify()
				notifyTemplateGroup.POST("/send-notify",
					casbinMiddleware.RequirePermission("system:notify-template:send-notify"),
					handlers.Notify.SendNotify)
			}

			// SMS Channel Protected Routes
			smsChannelProtectedGroup := systemGroup.Group("/sms-channel")
			{
				smsChannelProtectedGroup.POST("/create", casbinMiddleware.RequirePermission("system:sms-channel:create"), handlers.SmsChannel.CreateSmsChannel)
				smsChannelProtectedGroup.PUT("/update", casbinMiddleware.RequirePermission("system:sms-channel:update"), handlers.SmsChannel.UpdateSmsChannel)
				smsChannelProtectedGroup.DELETE("/delete", casbinMiddleware.RequirePermission("system:sms-channel:delete"), handlers.SmsChannel.DeleteSmsChannel)
				smsChannelProtectedGroup.GET("/page", casbinMiddleware.RequirePermission("system:sms-channel:query"), handlers.SmsChannel.GetSmsChannelPage)
				smsChannelProtectedGroup.GET("/get", casbinMiddleware.RequirePermission("system:sms-channel:query"), handlers.SmsChannel.GetSmsChannel)
				smsChannelProtectedGroup.GET("/simple-list", handlers.SmsChannel.GetSimpleSmsChannelList)
			}

			// SMS Template Protected Routes
			smsTemplateProtectedGroup := systemGroup.Group("/sms-template")
			{
				smsTemplateProtectedGroup.POST("/create", casbinMiddleware.RequirePermission("system:sms-template:create"), handlers.SmsTemplate.CreateSmsTemplate)
				smsTemplateProtectedGroup.PUT("/update", casbinMiddleware.RequirePermission("system:sms-template:update"), handlers.SmsTemplate.UpdateSmsTemplate)
				smsTemplateProtectedGroup.DELETE("/delete", casbinMiddleware.RequirePermission("system:sms-template:delete"), handlers.SmsTemplate.DeleteSmsTemplate)
				smsTemplateProtectedGroup.GET("/page", casbinMiddleware.RequirePermission("system:sms-template:query"), handlers.SmsTemplate.GetSmsTemplatePage)
				smsTemplateProtectedGroup.GET("/get", casbinMiddleware.RequirePermission("system:sms-template:query"), handlers.SmsTemplate.GetSmsTemplate)
				smsTemplateProtectedGroup.POST("/send", casbinMiddleware.RequirePermission("system:sms-template:send"), handlers.SmsTemplate.SendSms)
			}

			// SMS Log Protected Routes
			smsLogProtectedGroup := systemGroup.Group("/sms-log")
			{
				smsLogProtectedGroup.GET("/page", casbinMiddleware.RequirePermission("system:sms-log:query"), handlers.SmsLog.GetSmsLogPage)
				smsLogProtectedGroup.GET("/get", casbinMiddleware.RequirePermission("system:sms-log:query"), handlers.SmsLog.GetSmsLogPage)
			}

			// Log Protected Routes
			systemGroup.GET("/login-log/page", casbinMiddleware.RequirePermission("system:login-log:query"), handlers.LoginLog.GetLoginLogPage)
			systemGroup.GET("/operate-log/page", casbinMiddleware.RequirePermission("system:operate-log:query"), handlers.OperateLog.GetOperateLogPage)

			// File Protected Routes
			systemGroup.POST("/file/upload", infraHandlers.File.UploadFile)

			// Config Protected Routes
			configProtectedGroup := systemGroup.Group("/config")
			{
				configProtectedGroup.GET("/get", casbinMiddleware.RequirePermission("system:config:query"), infraHandlers.Config.GetConfig)
				configProtectedGroup.GET("/get-value-by-key", infraHandlers.Config.GetConfigKey) // Corrected method name
				configProtectedGroup.POST("/create", casbinMiddleware.RequirePermission("system:config:create"), infraHandlers.Config.CreateConfig)
				configProtectedGroup.PUT("/update", casbinMiddleware.RequirePermission("system:config:update"), infraHandlers.Config.UpdateConfig)
				configProtectedGroup.DELETE("/delete", casbinMiddleware.RequirePermission("system:config:delete"), infraHandlers.Config.DeleteConfig)
				configProtectedGroup.GET("/page", casbinMiddleware.RequirePermission("system:config:query"), infraHandlers.Config.GetConfigPage)
			}
		}

		// ====== Infra Routes (Public) ======
		infraPublicGroup := api.Group("/infra")
		{
			infraPublicGroup.GET("/file/:configId/get/*path", infraHandlers.File.GetFileContent)
		}

		// ====== Infra Routes (Protected) ======
		infraGroup := api.Group("/infra", middleware.Auth())
		{
			// WebSocket (对齐 Java /infra/ws)
			infraGroup.GET("/ws", infraHandlers.WebSocket.Handle)

			// File Config
			fileConfigGroup := infraGroup.Group("/file-config")
			{
				fileConfigGroup.POST("/create", casbinMiddleware.RequirePermission("infra:file-config:create"), infraHandlers.FileConfig.CreateFileConfig)
				fileConfigGroup.PUT("/update", casbinMiddleware.RequirePermission("infra:file-config:update"), infraHandlers.FileConfig.UpdateFileConfig)
				fileConfigGroup.PUT("/update-master", casbinMiddleware.RequirePermission("infra:file-config:update"), infraHandlers.FileConfig.UpdateFileConfigMaster)
				fileConfigGroup.DELETE("/delete", casbinMiddleware.RequirePermission("infra:file-config:delete"), infraHandlers.FileConfig.DeleteFileConfig)
				fileConfigGroup.GET("/page", casbinMiddleware.RequirePermission("infra:file-config:query"), infraHandlers.FileConfig.GetFileConfigPage)
				fileConfigGroup.GET("/get", casbinMiddleware.RequirePermission("infra:file-config:query"), infraHandlers.FileConfig.GetFileConfig)
				fileConfigGroup.GET("/test", casbinMiddleware.RequirePermission("infra:file-config:query"), infraHandlers.FileConfig.TestFileConfig)
			}

			// File
			fileGroup := infraGroup.Group("/file")
			{
				fileGroup.POST("/upload", infraHandlers.File.UploadFile)
				fileGroup.DELETE("/delete", casbinMiddleware.RequirePermission("infra:file:delete"), infraHandlers.File.DeleteFile)
				fileGroup.GET("/page", casbinMiddleware.RequirePermission("infra:file:query"), infraHandlers.File.GetFilePage)
				fileGroup.GET("/presigned-url", casbinMiddleware.RequirePermission("infra:file:query"), infraHandlers.File.GetFilePresignedUrl)
				fileGroup.POST("/create", casbinMiddleware.RequirePermission("infra:file:create"), infraHandlers.File.CreateFile)
				fileGroup.GET("/{config_id}/*path", infraHandlers.File.GetFileContent)
			}

			// Job
			jobGroup := infraGroup.Group("/job")
			{
				jobGroup.POST("/create", casbinMiddleware.RequirePermission("infra:job:create"), infraHandlers.Job.CreateJob)
				jobGroup.PUT("/update", casbinMiddleware.RequirePermission("infra:job:update"), infraHandlers.Job.UpdateJob)
				jobGroup.PUT("/update-status", casbinMiddleware.RequirePermission("infra:job:update"), infraHandlers.Job.UpdateJobStatus)
				jobGroup.DELETE("/delete", casbinMiddleware.RequirePermission("infra:job:delete"), infraHandlers.Job.DeleteJob)
				jobGroup.GET("/get", casbinMiddleware.RequirePermission("infra:job:query"), infraHandlers.Job.GetJob)
				jobGroup.GET("/page", casbinMiddleware.RequirePermission("infra:job:query"), infraHandlers.Job.GetJobPage)
				jobGroup.PUT("/trigger", casbinMiddleware.RequirePermission("infra:job:trigger"), infraHandlers.Job.TriggerJob)
				jobGroup.POST("/sync", casbinMiddleware.RequirePermission("infra:job:create"), infraHandlers.Job.SyncJob)
				jobGroup.GET("/export-excel", casbinMiddleware.RequirePermission("infra:job:export"), infraHandlers.Job.ExportJobExcel)
				jobGroup.GET("/get_next_times", casbinMiddleware.RequirePermission("infra:job:query"), infraHandlers.Job.GetJobNextTimes)
			}

			// Job Log
			jobLogGroup := infraGroup.Group("/job-log")
			{
				jobLogGroup.GET("/get", casbinMiddleware.RequirePermission("infra:job:query"), infraHandlers.JobLog.GetJobLog)
				jobLogGroup.GET("/page", casbinMiddleware.RequirePermission("infra:job:query"), infraHandlers.JobLog.GetJobLogPage)
				jobLogGroup.GET("/export-excel", casbinMiddleware.RequirePermission("infra:job:export"), infraHandlers.JobLog.ExportJobLogExcel)
			}

			// API Access Log
			apiAccessLogGroup := infraGroup.Group("/api-access-log")
			{
				apiAccessLogGroup.GET("/page", casbinMiddleware.RequirePermission("infra:api-access-log:query"), infraHandlers.ApiAccessLog.GetApiAccessLogPage)
			}

			// API Error Log
			apiErrorLogGroup := infraGroup.Group("/api-error-log")
			{
				apiErrorLogGroup.GET("/page", casbinMiddleware.RequirePermission("infra:api-error-log:query"), infraHandlers.ApiErrorLog.GetApiErrorLogPage)
				apiErrorLogGroup.PUT("/update-status", casbinMiddleware.RequirePermission("infra:api-error-log:update-status"), infraHandlers.ApiErrorLog.UpdateApiErrorLogProcess)
				apiErrorLogGroup.PUT("/update-process", casbinMiddleware.RequirePermission("infra:api-error-log:update-process"), infraHandlers.ApiErrorLog.UpdateApiErrorLogProcess)
			}
		}
	}
}

// RegisterAreaRoutes 注册地区路由 (Public - 不需要认证)
func RegisterAreaRoutes(engine *gin.Engine, areaHandler *system.AreaHandler) {
	api := engine.Group("/admin-api")
	{
		// Area 地区 (Public Routes)
		areaGroup := api.Group("/system/area")
		{
			areaGroup.GET("/tree", areaHandler.GetAreaTree)
			areaGroup.GET("/get-by-ip", areaHandler.GetAreaByIP)
		}
	}
}
