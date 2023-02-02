package logic

import (
	"context"
	"dining_home_backend_newest/service/main/internal/svc"
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

	print(url.ParseQuery(r.URL.RawQuery))

	return
}
