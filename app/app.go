package app

import (
	"context"

	"github.com/VQIVS/FlowForge/config"
	"github.com/VQIVS/FlowForge/internal/user"
	userPort "github.com/VQIVS/FlowForge/internal/user/port"
	"github.com/VQIVS/FlowForge/pkg/adapters/storage"
	appCtx "github.com/VQIVS/FlowForge/pkg/context"
	"github.com/VQIVS/FlowForge/pkg/postgres"
	"gorm.io/gorm"
)

type app struct {
	db          *gorm.DB
	cfg         config.Config
	userService userPort.Service
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) setDB() error {
	db, err := postgres.NewPostgormConnection(postgres.DBConnOptions{
		User:     a.cfg.DB.User,
		Password: a.cfg.DB.Password,
		Host:     a.cfg.DB.Host,
		DBName:   a.cfg.DB.Database,
		Schema:   a.cfg.DB.Schema,
	})
	if err != nil {
		return err
	}
	a.db = db
	return nil
}

func (a *app) UserService(ctx context.Context) userPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.userService == nil {
			a.userService = a.userServiceWithDB(a.db)
		}
		return a.userService
	}

	return a.userServiceWithDB(db)
}

func (a *app) userServiceWithDB(db *gorm.DB) userPort.Service {
	return user.NewService(storage.NewUserRepo(db))
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}
	return a, nil

}

func NewMustApp(cfg config.Config) App {
	a, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return a
}
