package handler

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/wcreate/tkits"
	"github.com/wcreate/wuc/api"
	"github.com/wcreate/wuc/models"
	"gopkg.in/macaron.v1"
)

// PUT /api/user/pwd/:uid/
func ModifyPassword(ctx *macaron.Context) {
	// 1.0
	var mpwd api.ModifyPasswordReq
	uid, ok := getUidAndBodyWithAuth(ctx, &mpwd)
	if !ok {
		return
	}

	// 2.0
	u := &models.User{Id: uid}
	if err := u.ReadOneOnly("Salt", "Password"); err == orm.ErrNoRows {
		ctx.JSON(404, api.INVALID_USER)
		return
	} else if err != nil {
		ctx.JSON(500, tkits.DB_ERROR)
		return
	}

	if !tkits.CmpPasswd(mpwd.OldPasswd, u.Salt, u.Password) {
		ctx.JSON(404, api.INVALID_USER)
		return
	}

	// 3.0
	pwd, salt := tkits.GenPasswd(mpwd.NewPasswd, 8)
	u.Salt = salt
	u.Password = pwd
	u.Updated = time.Now()

	if row, _ := u.Update("Salt", "Password", "Updated"); row != 1 {
		ctx.JSON(500, tkits.DB_ERROR)
		return
	}

	ctx.Status(200)
}
