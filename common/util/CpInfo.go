package util

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var (
	configAccessToken = "d:config:AccessToken"
	oncer             sync.Once
	Redis             *redis.Client
)

func InitRedisToken(rc *redis.Client) {
	oncer.Do(func() {
		Redis = rc
	})
}

func AccessTokenProvider(corpId string, corpSecret string, forceUpdate bool) string {
	ctx := context.Background()
	val, e := Redis.Get(ctx, configAccessToken).Result()
	if e == redis.Nil || forceUpdate {
		// 缓存中没有access key ，获取更新
		AccessToken, err := GetLatestAccessTokenToRedis(corpId, corpSecret)
		if err != nil {
			return ""
		}

		Redis.Set(ctx, configAccessToken, AccessToken, time.Minute*90)

		return AccessToken

	}
	return val

}
