package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	Mysql struct {
		Datasource string
	}

	Redis struct {
		Addr     string
		Password string
		DB       int
	}

	WxApp struct {
		CorpId    string
		AgentId   string
		AppSecret string
	}
}
