package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-macaron/captcha"
	"github.com/wcreate/tkits"
	"github.com/wcreate/wuc/api"
	"github.com/wcreate/wuc/models"
	"gopkg.in/macaron.v1"
)

// POST /api/user/login
func LoginUser(ctx *macaron.Context, cpt *captcha.Captcha) {
	var ulr api.UserLoginReq
	ok := getBody(ctx, &ulr)
	if !ok {
		return
	}

	if !cpt.Verify(ulr.CaptchaId, ulr.CaptchaName) {
		ctx.JSON(http.StatusBadRequest, api.INVALID_CAPTCHA)
		return
	}

	// check user whether existed
	u := &models.User{}
	if err := u.Find(ulr.Email, ulr.Username, ulr.Mobile); err != nil {
		ctx.JSON(http.StatusNotFound, api.INVALID_USER)
		return
	}

	// check user password
	if !tkits.CmpPasswd(ulr.Passwd, u.Salt, u.Password) {
		ctx.JSON(http.StatusNotFound, api.INVALID_USER)
		return
	}

	// update ip, time and count for login
	cip := getClientIP(ctx)
	u.LastLoginTime = time.Now()
	u.LastLoginIp = cip
	u.LoginCount += 1
	if _, err := u.Update("LastLoginTime", "LastLoginIp", "LoginCount"); err != nil {
		ctx.JSON(http.StatusInternalServerError, tkits.DB_ERROR)
		return
	}

	// generate a token
	suid := fmt.Sprintf("%v", u.Id)
	if token, err := tkits.GetSimpleToken().GenToken(cip, suid); err != nil {
		ctx.JSON(http.StatusInternalServerError, tkits.SYS_ERROR)
		return
	} else {
		rsp := &api.UserLoginRsp{}
		rsp.Uid = u.Id
		rsp.Username = u.Username
		rsp.Token = token

		ctx.SetCookie("token", token)
		ctx.SetCookie("uid", suid)
		
		ctx.JSON(http.StatusOK, rsp)
	}
}
