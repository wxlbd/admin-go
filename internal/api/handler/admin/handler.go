package admin

import (
	"github.com/wxlbd/admin-go/internal/api/handler/admin/infra"
	"github.com/wxlbd/admin-go/internal/api/handler/admin/system"
)

type AdminHandlers struct {
	Infra  *infra.Handlers
	System *system.Handlers
}
