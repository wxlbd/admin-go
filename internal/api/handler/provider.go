package handler

import (
	"github.com/google/wire"
	"github.com/wxlbd/admin-go/internal/api/handler/admin"
	adminInfra "github.com/wxlbd/admin-go/internal/api/handler/admin/infra"
	adminSystem "github.com/wxlbd/admin-go/internal/api/handler/admin/system"
)

var ProviderSet = wire.NewSet(
	// Admin System Providers Only
	adminSystem.ProviderSet,
	adminInfra.ProviderSet,
	wire.Struct(new(admin.AdminHandlers), "*"),
)
