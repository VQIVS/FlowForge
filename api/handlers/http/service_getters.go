package http

// import (
// 	"context"

// 	"github.com/VQIVS/FlowForge/api/service"
// 	"github.com/VQIVS/FlowForge/app"
// 	"github.com/VQIVS/FlowForge/config"
// )

// // user service transient instance handler
// func userServiceGetter(appContainer app.App, cfg config.ServerConfig) ServiceGetter[*service.UserService] {
// 	return func(ctx context.Context) *service.UserService {
// 		return service.NewUserService(appContainer.UserService(ctx),
// 			cfg.Secret, cfg.AuthExpMin, cfg.AuthRefreshMin)
// 	}
// }
