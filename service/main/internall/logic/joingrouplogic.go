package logic

import (
	"context"

	"dining_home_backend_newest/service/main/internall/svc"
	"dining_home_backend_newest/service/main/internall/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJoinGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinGroupLogic {
	return &JoinGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinGroupLogic) JoinGroup() (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
