package handler

import (
	"github.com/wcreate/tkits"
	"github.com/wcreate/wuc/rest"
	"gopkg.in/macaron.v1"

	"net/http"
)

// Get /api/user/info/:uid/
func UserInfo(ctx *macaron.Context, as rest.AuthService, ut *rest.UserToken) {
	// 1.0
	uid, ok := getUidWithAuth(ctx, as, ut, rest.DummyOptId)
	if !ok {
		return
	}

	rsp := &rest.UserInfoRsp{}
	rsp.Id = uid
	if err := rsp.QueryAll(); err != nil {
		ctx.JSON(http.StatusInternalServerError, tkits.DB_ERROR)
		return
	}

	ctx.JSON(http.StatusOK, rsp)
}
