package svc

import (
	"dining_home_backend_newest/service/main/internal/config"
	"dining_home_backend_newest/service/main/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config           config.Config
	GlobalMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		GlobalMiddleware: middleware.NewGlobalMiddleware().Handle,
	}
}
