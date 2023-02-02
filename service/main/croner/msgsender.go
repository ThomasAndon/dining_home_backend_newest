package croner

import (
	"dining_home_backend_newest/common/util"
	"github.com/ricardolonga/jsongo"
	"github.com/tidwall/gjson"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"strings"
)

const (
	success = iota
	token_expired
	failed
)

func SendMsg(toUserId []string, title, content, url string) {
	userIds := strings.Join(toUserId, "|")
	reqBody := jsongo.Object().Put("touser", userIds).Put("msgtype", "textcard").
		Put("agentid", 1000002).Put("textcard", jsongo.Object().Put("title", title).
		Put("description", content).Put("url", url).Put("btntxt", "点击查看详情")).
		Put("enable_id_trans", 0).Put("enable_duplicate_check", 0).
		Put("duplicate_check_interval", 1800)

	//logx.Info(reqBody.String())

	res := sendRequest(reqBody.String(), false)

	if res == token_expired {
		sendRequest(reqBody.String(), true)
	}

	//preparedUrl := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" +
	//	util.AccessTokenProvider(false)

}

func sendRequest(bodyJsonString string, force bool) int {
	preparedUrl := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" +
		util.AccessTokenProvider(force)
	body := strings.NewReader(bodyJsonString)
	req, _ := http.NewRequest("POST", preparedUrl, body)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		return failed
	}
	respBody, _ := io.ReadAll(resp.Body)
	logx.Info(string(respBody))
	errCode := gjson.Get(string(respBody), "errcode").Int()
	if errCode == 42001 {
		return token_expired
	}

	if errCode == 0 {
		return success
	}

	return failed
}
