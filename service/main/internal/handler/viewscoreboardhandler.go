package handler

import (
	"net/http"

	"dining_home_backend_newest/service/main/internal/logic"
	"dining_home_backend_newest/service/main/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ViewScoreboardHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewViewScoreboardLogic(r.Context(), svcCtx)
		resp, err := l.ViewScoreboard()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
