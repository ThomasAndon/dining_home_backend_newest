package main

import (
	"dining_home_backend_newest/service/main/croner"
	"flag"
	"fmt"

	"dining_home_backend_newest/service/main/internal/config"
	"dining_home_backend_newest/service/main/internal/handler"
	"dining_home_backend_newest/service/main/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/main.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	go croner.CronJob(ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()

}
