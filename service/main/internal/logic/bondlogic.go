package logic

import (
	"context"
	"dining_home_backend_newest/service/main/internal/svc"
	"dining_home_backend_newest/service/main/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BondLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBondLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BondLogic {
	return &BondLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BondLogic) Bond(req *types.BondRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	a := l.ctx.Value("mwKey")
	print(a)
	return
}
