package svc

import (
	"dining_home_backend_newest/service/main/internal/config"
	"dining_home_backend_newest/service/main/internal/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config           config.Config
	GlobalMiddleware rest.Middleware
	redis            *redis.Client
	xorm             *xorm.Engine
}

func NewServiceContext(c config.Config) *ServiceContext {
	xEng, err := xorm.NewEngine("mysql", c.Mysql.Datasource)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:           c,
		GlobalMiddleware: middleware.NewGlobalMiddleware().Handle,
		redis: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Addr,
			Password: c.Redis.Password,
			DB:       c.Redis.DB,
		}),
		xorm: xEng,
	}
}
