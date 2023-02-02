package util

import (
	"context"
	"dining_home_backend_newest/service/main/internall/svc"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var (
	configAccessToken = "d:config:AccessToken"
	oncer             sync.Once
	//Redis             *redis.Client
	service *svc.ServiceContext
)

func InitRedisToken(svcCtx *svc.ServiceContext) {
	oncer.Do(func() {
		service = svcCtx
	})
}

func AccessTokenProvider(corpId string, corpSecret string, forceUpdate bool) string {
	ctx := context.Background()
	val, e := service.Redis.Get(ctx, configAccessToken).Result()
	if e == redis.Nil || forceUpdate {
		// 缓存中没有access key ，获取更新
		AccessToken, err := GetLatestAccessTokenToRedis(corpId, corpSecret)
		if err != nil {
			return ""
		}

		service.Redis.Set(ctx, configAccessToken, AccessToken, time.Minute*90)

		return AccessToken

	}
	return val

}
