package handler

import (
	"net/http"

	"dining_home_backend_newest/service/main/internal/logic"
	"dining_home_backend_newest/service/main/internal/svc"
	"dining_home_backend_newest/service/main/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BondHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BondRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewBondLogic(r.Context(), svcCtx)
		//headers := r.Header
		resp, err := l.Bond(&req, r.Header)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
