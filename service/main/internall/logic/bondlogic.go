package logic

import (
	"context"
	"dining_home_backend_newest/service/main/croner"
	"dining_home_backend_newest/service/main/internall/svc"
	"dining_home_backend_newest/service/main/internall/types"
	"net/http"

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

func (l *BondLogic) Bond(req *types.BondRequest, header http.Header) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	logx.Info("called")

	croner.SendMsg([]string{"TangCaitong"}, "请发送今日吃饭情况", "输入框中直接输入并发送：1：在家，2：在外，3：其他不准时回家吃饭的情况（点击查看家庭成员今日饭况）", "https://www.tando.fun/d")

	return
}
