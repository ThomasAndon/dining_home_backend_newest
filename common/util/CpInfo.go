package util

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	configAccessToken = "d:config:AccessToken"
)

func AccessTokenProvider(ctx context.Context, rc *redis.Client, corpId string, corpSecret string, forceUpdate bool) string {
	val, e := rc.Get(ctx, configAccessToken).Result()
	if e == redis.Nil || forceUpdate {
		// 缓存中没有access key ，获取更新
		AccessToken, err := getLatestAccessTokenToRedis(corpId, corpSecret)
		if err != nil {
			return ""
		}

		rc.Set(ctx, configAccessToken, AccessToken, time.Minute*90)

		return AccessToken

	}
	return val

}
