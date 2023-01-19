package handler

import (
	"net/http"

	"dining_home_backend_newest/service/main/internal/logic"
	"dining_home_backend_newest/service/main/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func JoinGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewJoinGroupLogic(r.Context(), svcCtx)
		resp, err := l.JoinGroup()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
