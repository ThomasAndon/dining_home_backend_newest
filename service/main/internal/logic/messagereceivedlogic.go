package logic

import (
	"context"

	"dining_home_backend_newest/service/main/internal/svc"
	"dining_home_backend_newest/service/main/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *MessageReceivedLogic) MessageReceived() (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
