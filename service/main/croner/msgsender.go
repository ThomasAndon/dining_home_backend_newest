package croner

import (
	"dining_home_backend_newest/common/util"
	"github.com/ricardolonga/jsongo"
	"strings"
)

func SendMsg(toUserId []string, title, content, url string) {
	userIds := strings.Join(toUserId, "|")
	reqBody := jsongo.Object().Put("touser", userIds).Put("msgtype", "textcard").
		Put("agentid", 1000002).Put("textcard", jsongo.Object().Put("title", title).
		Put("description", content).Put("url", url).Put("btntxt", "点击查看详情")).
		Put("enable_id_trans", 0).Put("enable_duplicate_check", 0).
		Put("duplicate_check_interval", 1800)

	//logx.Info(reqBody.String())
	preparedUrl := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" +
		util.AccessTokenProvider()

}
