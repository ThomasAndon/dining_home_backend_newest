package logic

import (
	"context"
	"dining_home_backend_newest/common/wxbizmsgcrypt"
	"dining_home_backend_newest/service/main/internal/svc"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"net/url"
)

type MessageReceivedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
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
	//body, e := io.ReadAll(r.Body)
	//if e != nil {
	//	return "", e
	//}
	res, _ := url.ParseQuery(r.URL.RawQuery)
	//logx.Info("hahaha ", res.Get("msg_signature"))
	//logx.Info("hahaha ", res.Get("timestamp"))
	//logx.Info("hahaha ", res.Get("nonce"))
	//logx.Info("hahaha ", res.Get("echostr"))
	token := "ohgl2y1Bpt5KWa14Fa62FufJhE"
	receiverId := "wwfc9c5ccc385b29d0"
	encodingAeskey := "QawdQh7xKlzIrwqhDD84tfUvBRvwyXdBY1V5zTWK8Bh"
	wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(token, encodingAeskey, receiverId, wxbizmsgcrypt.XmlType)

	verifyMsgSign := res.Get("msg_signature")
	verifyTimestamp := res.Get("timestamp")
	verifyNonce := res.Get("nonce")
	verifyEchoStr := res.Get("echostr")
	echoStr, cryptErr := wxcpt.VerifyURL(verifyMsgSign, verifyTimestamp, verifyNonce, verifyEchoStr)
	if nil != cryptErr {
		fmt.Println("verifyUrl fail", cryptErr)
	}
	fmt.Println("verifyUrl success echoStr", string(echoStr))
	resp = string(echoStr)
	return
}
