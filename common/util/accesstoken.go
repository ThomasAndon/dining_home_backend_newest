package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type accessTokenResp struct {
	errCode     int    `json:"errcode"`
	errmsg      string `json:"errmsg"`
	accessToken string `json:"access_token"`
	expiresIn   int    `json:"expires_in"`
}

func getLatestAccessTokenToRedis(corpId string, appSecret string) (string, error) {
	a := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?"+
		"corpid=%s&corpsecret=%s", corpId, appSecret)
	res, err := http.Get(a)

	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var resp accessTokenResp
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	json.Unmarshal(body, &resp)

	return resp.accessToken, nil
}
