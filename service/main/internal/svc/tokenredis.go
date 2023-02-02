package svc

import (
	"context"
	"dining_home_backend_newest/common/util"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	configAccessToken = "d:config:AccessToken"
)

func AccessTokenProvider(rc *redis.Client, corpId string, corpSecret string, forceUpdate bool) string {
	ctx := context.Background()
	val, e := rc.Get(ctx, configAccessToken).Result()
	if e == redis.Nil || forceUpdate {
		// 缓存中没有access key ，获取更新
		AccessToken, err := util.GetLatestAccessTokenToRedis(corpId, corpSecret)
		if err != nil {
			return ""
		}

		rc.Set(ctx, configAccessToken, AccessToken, time.Minute*90)

		return AccessToken

	}
	return val

}
