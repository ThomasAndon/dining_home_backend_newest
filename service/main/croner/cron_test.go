package croner

import (
	"fmt"
	"github.com/ricardolonga/jsongo"
	"github.com/tidwall/gjson"
	"testing"
)

func TestJsongo(t *testing.T) {
	userIds := "TangCaitong"
	content := "Content"
	url := "https://www.tando.fun/d"
	title := "title"
	reqBody := jsongo.Object().Put("touser", userIds).Put("msgtype", "textcard").
		Put("agentid", 1000002).Put("textcard", jsongo.Object().Put("title", title).
		Put("description", content).Put("url", url).Put("btntxt", "点击查看详情")).
		Put("enable_id_trans", 0).Put("enable_duplicate_check", 0).
		Put("duplicate_check_interval", 1800)

	fmt.Println(reqBody.Indent())
}

func TestDecode(t *testing.T) {
	resStr := `{
  "errcode" : 0,
  "errmsg" : "ok",
  "invaliduser" : "userid1|userid2",
  "invalidparty" : "partyid1|partyid2",
  "invalidtag": "tagid1|tagid2",
  "unlicenseduser" : "userid3|userid4",
  "msgid": "xxxx",
  "response_code": "xyzxyz"
}`
	res := gjson.Get(resStr, "errcode")
	t.Log(res.Int() == 0)
	// 42001 expire
}
