package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	JwtAuthMW struct {
		AccessSecret string
		AccessExpire int64
	}
}
