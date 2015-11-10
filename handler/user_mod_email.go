package handler

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/wcreate/tkits"
	"github.com/wcreate/wuc/models"
	"github.com/wcreate/wuc/rest"
	"gopkg.in/macaron.v1"
)

// PUT /api/user/email/:uid/
func ModifyEmail(ctx *macaron.Context, as rest.AuthService, ut *rest.UserToken) {
	// 1.0
	var memail rest.ModifyEmailReq
	uid, ok := getUidAndBodyWithAuth(ctx, as, ut, rest.DummyOptId, &memail)
	if !ok {
		return
	}

	// 2.0
	u := &models.User{Id: uid}
	if err := u.ReadOneOnly("Email"); err == orm.ErrNoRows {
		ctx.JSON(http.StatusNotFound, rest.INVALID_USER)
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, tkits.DB_ERROR)
		return
	}

	if memail.OldEmail != u.Email {
		ctx.JSON(http.StatusBadRequest, rest.INVALID_EMAIL)
		return
	}

	if !validReq(ctx, &memail) {
		return
	}

	// 3.0
	u.Email = memail.NewEmail
	u.Cfmcode = tkits.StringNewRand(32)

	if row, _ := u.Update("Email", "Cfmcode"); row != 1 {
		ctx.JSON(http.StatusInternalServerError, tkits.DB_ERROR)
		return
	}

	if err := SendCfmEmail(u, CFM_MOD_SUBJET); err != nil {
		ctx.JSON(http.StatusInternalServerError, rest.SEND_EMAIL_FAILED)
		return
	}

	rsp := &rest.ModifyEmailRsp{
		uid,
		fmt.Sprintf("http://%s//api/user/cfm?t=mail&amp;uid=%v&amp;code=%s", tkits.WebDomain, uid, u.Cfmcode),
	}

	ctx.JSON(http.StatusOK, rsp)
}
