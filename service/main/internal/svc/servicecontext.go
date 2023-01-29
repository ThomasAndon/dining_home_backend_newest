package svc

import (
	"dining_home_backend_newest/service/main/internal/config"
	"dining_home_backend_newest/service/main/internal/middleware"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config           config.Config
	GlobalMiddleware rest.Middleware
	redis            *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		GlobalMiddleware: middleware.NewGlobalMiddleware().Handle,
		redis: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Addr,
			Password: c.Redis.Password,
			DB:       c.Redis.DB,
		}),
	}
}
