package handler

import (
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/wcreate/tkits"
	"github.com/wcreate/wuc/models"
	"github.com/wcreate/wuc/rest"
	"gopkg.in/macaron.v1"
)

// POST/PUT /api/user/info/:uid/
func ModifyUser(ctx *macaron.Context, as rest.AuthService, ut *rest.UserToken) {
	// 1.0
	var mui rest.ModifyUserInfoReq
	uid, ok := getUidAndBodyWithAuth(ctx, as, ut, rest.DummyOptId, &mui)
	if !ok {
		return
	}

	if uid != mui.Uid {
		ctx.JSON(http.StatusBadRequest, rest.INVALID_USER)
		return
	}
	mui.User = &models.User{Id: uid}

	// 1.1 TODO Check the params
	if !validReq(ctx, &mui) {
		return
	}

	// 2.0
	oldui := &models.UserInfo{User: &models.User{Id: uid}}
	if err := oldui.Read("User"); err == orm.ErrNoRows {
		// not exist then insert it
		if _, err := mui.Insert(); err != nil {
			ctx.JSON(http.StatusInternalServerError, tkits.DB_ERROR)
			return
		}
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, tkits.DB_ERROR)
		return
	}

	// 3.0
	if err := mui.UpdateInfo(); err != nil {
		ctx.JSON(http.StatusInternalServerError, tkits.DB_ERROR)
		return
	}

	ctx.Status(http.StatusOK)
}
