package app

import (
	"context"

	"github.com/VQIVS/FlowForge/config"
	userPort "github.com/VQIVS/FlowForge/internal/user/port"
	"gorm.io/gorm"
)

type App interface {
	UserService(ctx context.Context) userPort.Service
	DB() *gorm.DB
	Config() config.Config
}
