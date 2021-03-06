package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-macaron/captcha"
	"github.com/wcreate/tkits"
	"github.com/wcreate/wuc/models"
	"github.com/wcreate/wuc/rest"
	"gopkg.in/macaron.v1"
)

// POST /api/user/login
func LoginUser(ctx *macaron.Context, as rest.AuthService, cpt *captcha.Captcha) {
	var ulr rest.UserLoginReq
	ok := getBody(ctx, &ulr)
	if !ok {
		return
	}

	if !cpt.Verify(ulr.CaptchaId, ulr.CaptchaValue) {
		ctx.JSON(http.StatusBadRequest, rest.INVALID_CAPTCHA)
		return
	}

	// check user whether existed
	u := &models.User{}
	if err := u.Find(ulr.Email, ulr.Username, ulr.Mobile); err != nil {
		ctx.JSON(http.StatusNotFound, rest.INVALID_USER)
		return
	}

	// check user password
	if !tkits.CmpPasswd(ulr.Passwd, u.Salt, u.Password) {
		ctx.JSON(http.StatusNotFound, rest.INVALID_USER)
		return
	}

	// update ip, time and count for login
	cip := ctx.RemoteAddr()
	u.LastLoginTime = time.Now()
	u.LastLoginIp = cip
	u.LoginCount += 1
	if _, err := u.Update("LastLoginTime", "LastLoginIp", "LoginCount"); err != nil {
		ctx.JSON(http.StatusInternalServerError, tkits.DB_ERROR)
		return
	}

	// generate a token

	if token, err := as.GenUserToken(cip, u.Id, 15, rest.TokenUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, tkits.SYS_ERROR)
		return
	} else {
		rsp := &rest.UserLoginRsp{}
		rsp.Uid = u.Id
		rsp.Username = u.Username
		rsp.Token = token

		if ulr.CookieMaxAge == 0 {
			ulr.CookieMaxAge = 60 * 60 * 12 //half of one day
		}

		suid := fmt.Sprintf("%v", u.Id)
		ctx.SetCookie("token", token, ulr.CookieMaxAge)
		ctx.SetCookie("uid", suid, ulr.CookieMaxAge)

		ctx.JSON(http.StatusOK, rsp)
	}
}
