package handler

import (
	"net/http"

	"dining_home_backend_newest/service/main/internall/logic"
	"dining_home_backend_newest/service/main/internall/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUpdateStatusLogic(r.Context(), svcCtx)
		resp, err := l.UpdateStatus()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
