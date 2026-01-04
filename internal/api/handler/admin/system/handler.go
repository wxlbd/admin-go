package system

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewAreaHandler,
	NewAuthHandler,
	NewDeptHandler,
	NewDictHandler,
	NewLoginLogHandler,
	NewMenuHandler,
	NewNoticeHandler,
	NewNotifyHandler,
	NewOperateLogHandler,
	NewPermissionHandler,
	NewPostHandler,
	NewRoleHandler,
	NewTenantHandler,
	NewTenantPackageHandler,
	NewUserHandler,
	NewSmsChannelHandler,
	NewSmsTemplateHandler,
	NewSmsLogHandler,
	NewHandlers,
)

type Handlers struct {
	Area          *AreaHandler
	Auth          *AuthHandler
	Dept          *DeptHandler
	Dict          *DictHandler
	LoginLog      *LoginLogHandler
	Menu          *MenuHandler
	Notice        *NoticeHandler
	Notify        *NotifyHandler
	OperateLog    *OperateLogHandler
	Permission    *PermissionHandler
	Post          *PostHandler
	Role          *RoleHandler
	Tenant        *TenantHandler
	TenantPackage *TenantPackageHandler
	User          *UserHandler
	SmsChannel    *SmsChannelHandler
	SmsTemplate   *SmsTemplateHandler
	SmsLog        *SmsLogHandler
}

func NewHandlers(
	area *AreaHandler,
	auth *AuthHandler,
	dept *DeptHandler,
	dict *DictHandler,
	loginLog *LoginLogHandler,
	menu *MenuHandler,
	notice *NoticeHandler,
	notify *NotifyHandler,
	operateLog *OperateLogHandler,
	permission *PermissionHandler,
	post *PostHandler,
	role *RoleHandler,
	tenant *TenantHandler,
	tenantPackage *TenantPackageHandler,
	user *UserHandler,
	smsChannel *SmsChannelHandler,
	smsTemplate *SmsTemplateHandler,
	smsLog *SmsLogHandler,
) *Handlers {
	return &Handlers{
		Area:          area,
		Auth:          auth,
		Dept:          dept,
		Dict:          dict,
		LoginLog:      loginLog,
		Menu:          menu,
		Notice:        notice,
		Notify:        notify,
		OperateLog:    operateLog,
		Permission:    permission,
		Post:          post,
		Role:          role,
		Tenant:        tenant,
		TenantPackage: tenantPackage,
		User:          user,
		SmsChannel:    smsChannel,
		SmsTemplate:   smsTemplate,
		SmsLog:        smsLog,
	}
}
