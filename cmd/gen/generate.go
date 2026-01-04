package main

import (
	"github.com/wxlbd/admin-go/internal/model"

	"gorm.io/gen"
)

func main() {
	// 1. 不需要连接数据库，直接基于 Struct 生成

	// 2. 配置生成器
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./internal/repo/query",
		ModelPkgPath:  "./internal/model",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	// 3. 基于 Struct 生成 - 仅系统管理模型
	g.ApplyBasic(
		// System
		model.SystemUser{},
		model.SystemRole{},
		model.SystemMenu{},
		model.SystemTenant{},
		model.SystemDictData{},
		model.SystemDictType{},
		model.SystemDept{},
		model.SystemPost{},
		model.SystemUserPost{},
		model.SystemNotice{},
		model.SystemUserRole{},
		model.SystemRoleMenu{},
		model.SystemConfig{},
		model.SystemSmsChannel{},
		model.SystemSmsTemplate{},
		model.SystemSmsLog{},
		model.SystemSmsCode{},
		model.InfraFileConfig{},
		model.InfraFile{},
		model.SocialUser{},
		model.SocialUserBind{},
		model.SocialClient{},
		model.SystemLoginLog{},
		model.SystemOperateLog{},
		&model.SystemNotifyMessage{},
		&model.SystemNotifyTemplate{},
		model.InfraJob{},
		model.InfraJobLog{},
		model.InfraApiAccessLog{},
		model.InfraApiErrorLog{},
		model.SystemTenantPackage{},
	)

	// 4. 执行生成
	g.Execute()
}
