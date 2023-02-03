package croner

import (
	"context"
	"dining_home_backend_newest/common/util"
	"dining_home_backend_newest/service/main/internall/svc"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
	"sync"
)

var (
	tool  *svc.ServiceContext
	oncer sync.Once
)

func CronJob(svcCtx *svc.ServiceContext) {
	oncer.Do(func() {
		tool = svcCtx
	})
	c := cron.New()
	c.AddFunc("9 3 17 * * *", doTodayWork)
	c.Start()
	InitCronContext(svcCtx)
	go redisListen(svcCtx.Redis)

}

func InitCronContext(svcCtx *svc.ServiceContext) {
	util.InitRedisToken(svcCtx)
}

func doTodayWork() {

}

func redisListen(r *redis.Client) {
	ctx := context.Background()
	pubsub := r.Subscribe(ctx, "__keyevent@0__:expired")
	for {
		msg, _ := pubsub.ReceiveMessage(ctx)
		//if err != nil {
		//	panic(err)
		//}

		redisCalledBack(msg.Payload)
	}
}
