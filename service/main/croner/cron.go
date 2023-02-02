package croner

import (
	"context"
	"dining_home_backend_newest/common/util"
	"dining_home_backend_newest/service/main/internal/svc"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
)

func CronJob(svcCtx *svc.ServiceContext) {
	c := cron.New()
	c.AddFunc("0 3 17 * * *", doTodaysWork)
	InitContext(svcCtx)
	go redisListen(svcCtx.Redis)

}

func InitContext(svcCtx *svc.ServiceContext) {
	util.InitRedisToken(svcCtx.Redis)
}

func doTodaysWork() {

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
