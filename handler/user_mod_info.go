package handler

import (
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/wcreate/tkits"
	"github.com/wcreate/wuc/api"
	"github.com/wcreate/wuc/models"
	"gopkg.in/macaron.v1"
)

// POST/PUT /api/user/info/:uid/
func ModifyUser(ctx *macaron.Context) {
	// 1.0
	var mui api.ModifyUserInfoReq
	uid, ok := getUidAndBodyWithAuth(ctx, &mui)
	if !ok {
		return
	}

	if uid != mui.Uid {
		ctx.JSON(http.StatusBadRequest, api.INVALID_USER)
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
