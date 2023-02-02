package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type accessTokenResp struct {
	ErrCode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func getLatestAccessTokenToRedis(corpId string, appSecret string) (string, error) {
	a := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?"+
		"corpid=%s&corpsecret=%s", corpId, appSecret)
	res, err := http.Get(a)

	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	resp := accessTokenResp{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	//buf := new(strings.Builder)
	//n, err := io.Copy(buf, res.Body)

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return "", err
	}

	return resp.AccessToken, nil
}
