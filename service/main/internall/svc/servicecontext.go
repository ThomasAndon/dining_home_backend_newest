package svc

import (
	"dining_home_backend_newest/service/main/internall/config"
	"dining_home_backend_newest/service/main/internall/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"xorm.io/core"
)

type ServiceContext struct {
	Config           config.Config
	GlobalMiddleware rest.Middleware
	Redis            *redis.Client
	Xorm             *xorm.Engine
}

func NewServiceContext(c config.Config) *ServiceContext {
	xEng, err := xorm.NewEngine("mysql", c.Mysql.Datasource)
	xEng.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, "d_"))
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:           c,
		GlobalMiddleware: middleware.NewGlobalMiddleware().Handle,
		Redis: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Addr,
			Password: c.Redis.Password,
			DB:       c.Redis.DB,
		}),
		Xorm: xEng,
	}
}
