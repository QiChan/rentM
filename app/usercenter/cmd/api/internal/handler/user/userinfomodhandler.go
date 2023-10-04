package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"rentM/app/usercenter/cmd/api/internal/logic/user"
	"rentM/app/usercenter/cmd/api/internal/svc"
	"rentM/app/usercenter/cmd/api/internal/types"
)

func UserInfoModHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoModReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserInfoModLogic(r.Context(), svcCtx)
		resp, err := l.UserInfoMod(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
