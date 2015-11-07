package handler

import (
	"fmt"
	"time"

	"github.com/go-macaron/captcha"
	"github.com/wcreate/wuc/api"
	"github.com/wcreate/wuc/models"
	"github.com/wcreate/wuc/security"
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
		ctx.JSON(400, api.INVALID_CAPTCHA)
		return
	}

	// check user whether existed
	u := models.User{}
	if err := u.Find(ulr.Email, ulr.Username, ulr.Mobile); err != nil {
		ctx.JSON(400, api.INVALID_USER)
		return
	}

	// check user password
	if !security.CmpPasswd(ulr.Passwd, u.Salt, u.Password) {
		ctx.JSON(404, api.INVALID_USER)
		return
	}

	// update ip, time and count for login
	cip := getClientIP(ctx)
	u.LastLoginTime = time.Now()
	u.LastLoginIp = cip
	u.LoginCount += 1
	if _, err := u.Update("LastLoginTime", "LastLoginIp", "LoginCount"); err != nil {
		ctx.JSON(400, api.DB_ERROR)
		return
	}

	// generate a token
	if token, err := security.GetDefaultSimpleToken().GenToken(cip, fmt.Sprintf("%v", u.Id)); err != nil {
		ctx.JSON(404, api.SYS_ERROR)
		return
	} else {
		rsp := &api.UserLoginRsp{
			u.Id,
			u.Username,
			token,
		}
		ctx.JSON(200, rsp)
	}
}
