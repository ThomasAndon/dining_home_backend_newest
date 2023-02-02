package logic

import (
	"context"
	"dining_home_backend_newest/common/wxbizmsgcrypt"
	"dining_home_backend_newest/service/main/internall/svc"
	models "dining_home_backend_newest/service/main/model"
	"encoding/xml"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"net/url"
)

type MessageReceivedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}
type MsgContent struct {
	ToUsername   string `xml:"ToUserName"`
	FromUsername string `xml:"FromUserName"`
	CreateTime   uint32 `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	Msgid        string `xml:"MsgId"`
	Agentid      uint32 `xml:"AgentId"`
}

func NewMessageReceivedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageReceivedLogic {
	return &MessageReceivedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageReceivedLogic) MessageReceived(r *http.Request) (resp string, err error) {
	// todo: add your logic here and delete this line
	body, e := io.ReadAll(r.Body)
	if e != nil {
		return "", e
	}
	//print(body)
	parseResult, _ := url.ParseQuery(r.URL.RawQuery)

	token := "ohgl2y1Bpt5KWa14Fa62FufJhE"
	receiverId := "wwfc9c5ccc385b29d0"
	encodingAeskey := "QawdQh7xKlzIrwqhDD84tfUvBRvwyXdBY1V5zTWK8Bh"
	wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(token, encodingAeskey, receiverId, wxbizmsgcrypt.XmlType)

	reqMsgSign := parseResult.Get("msg_signature")
	reqTimestamp := parseResult.Get("timestamp")
	reqNonce := parseResult.Get("nonce")

	reqData := body
	msg, cryptErr := wxcpt.DecryptMsg(reqMsgSign, reqTimestamp, reqNonce, reqData)
	if nil != cryptErr {
		fmt.Println("verifyUrl fail", cryptErr)
	}
	fmt.Println("after decrypt msg: ", string(msg))

	var msgContent MsgContent
	err = xml.Unmarshal(msg, &msgContent)
	if nil != err {
		fmt.Println("Unmarshal fail")
	} else {
		fmt.Println("struct", msgContent)
	}
	// todo
	person := models.Identity{CpwxId: msgContent.FromUsername}
	l.svcCtx.Xorm.Get(person)

	return
}
