package context

import (
	"context"
	"log/slog"

	"gorm.io/gorm"
)

var defaultLogger *slog.Logger

type appContext struct {
	context.Context
	db           *gorm.DB
	shouldCommit bool
	logger       *slog.Logger
}

type AppContextOpt func(*appContext) *appContext

func WithDB(db *gorm.DB, shouldCommit bool) AppContextOpt {
	return func(ac *appContext) *appContext {
		ac.db = db
		ac.shouldCommit = shouldCommit
		return ac
	}
}

func WithLogger(logger *slog.Logger) AppContextOpt {
	return func(ac *appContext) *appContext {
		ac.logger = logger
		return ac
	}
}

func SetDB(ctx context.Context, db *gorm.DB, shouldCommit bool) {
	appCtx, ok := ctx.(*appContext)
	if !ok {
		return
	}
	appCtx.db = db
	appCtx.shouldCommit = shouldCommit
}

func GetDB(ctx context.Context) *gorm.DB {
	appCtx, ok := ctx.(*appContext)
	if !ok {
		return nil
	}
	return appCtx.db
}

func Commit(ctx context.Context) error {
	appCtx, ok := ctx.(*appContext)
	if !ok || !appCtx.shouldCommit {
		return nil
	}
	return appCtx.db.Commit().Error
}

func Rollback(ctx context.Context) error {
	appCtx, ok := ctx.(*appContext)
	if !ok || !appCtx.shouldCommit {
		return nil
	}
	return appCtx.db.Rollback().Error
}

func SetLogger(ctx context.Context, logger *slog.Logger) {
	if appCtx, ok := ctx.(*appContext); ok {
		appCtx.logger = logger
	}
}

func GetLogger(ctx context.Context) *slog.Logger {
	appCtx, ok := ctx.(*appContext)
	if !ok || appCtx.logger == nil {

		return defaultLogger
	}
	return appCtx.logger
}

func NewAppContext(parent context.Context, opts ...AppContextOpt) context.Context {
	ctx := &appContext{Context: parent}
	for _, opt := range opts {
		ctx = opt(ctx)
	}

	return ctx
}
