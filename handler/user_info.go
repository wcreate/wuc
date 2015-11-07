package handler

import (
	"github.com/wcreate/tkits"
	"github.com/wcreate/wuc/api"
	"gopkg.in/macaron.v1"
)

// Get /api/user/info/:uid/
func UserInfo(ctx *macaron.Context) {
	// 1.0
	uid, ok := getUidWithAuth(ctx)
	if !ok {
		return
	}

	rsp := &api.UserInfoRsp{}
	rsp.Id = uid
	if err := rsp.QueryAll(); err != nil {
		ctx.JSON(500, tkits.DB_ERROR)
		return
	}

	ctx.JSON(200, rsp)
}
