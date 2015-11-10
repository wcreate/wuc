package handler

import (
	"net/http"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/wcreate/tkits"
	"github.com/wcreate/wuc/models"
	"github.com/wcreate/wuc/rest"
	"gopkg.in/macaron.v1"
)

// PUT /api/user/pwd/:uid/
func ModifyPassword(ctx *macaron.Context, as rest.AuthService, ut *rest.UserToken) {
	// 1.0
	var mpwd rest.ModifyPasswordReq
	uid, ok := getUidAndBodyWithAuth(ctx, as, ut, rest.DummyOptId, &mpwd)
	if !ok {
		return
	}

	// 2.0
	u := &models.User{Id: uid}
	if err := u.ReadOneOnly("Salt", "Password"); err == orm.ErrNoRows {
		ctx.JSON(http.StatusNotFound, rest.INVALID_USER)
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, tkits.DB_ERROR)
		return
	}

	if !tkits.CmpPasswd(mpwd.OldPasswd, u.Salt, u.Password) {
		ctx.JSON(http.StatusNotFound, rest.INVALID_USER)
		return
	}

	valid := validation.Validation{}
	valid.Match(mpwd.NewPasswd, rest.ValidPasswd,
		"NewPasswd").Message(rest.PasswdPrompt)
	if !validMember(ctx, &valid) {
		return
	}

	// 3.0
	pwd, salt := tkits.GenPasswd(mpwd.NewPasswd, 8)
	u.Salt = salt
	u.Password = pwd
	u.Updated = time.Now()

	if row, _ := u.Update("Salt", "Password", "Updated"); row != 1 {
		ctx.JSON(http.StatusInternalServerError, tkits.DB_ERROR)
		return
	}

	ctx.Status(http.StatusOK)
}
