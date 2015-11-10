package handler

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/go-macaron/captcha"
	"github.com/wcreate/tkits"
	"github.com/wcreate/wuc/models"
	"github.com/wcreate/wuc/rest"
	"gopkg.in/macaron.v1"
)

// POST /api/user/signup
func AddUser(ctx *macaron.Context, as rest.AuthService, cpt *captcha.Captcha) {
	var uar rest.UserAddReq
	ok := getBody(ctx, &uar)
	if !ok {
		return
	}

	log.Debugf("retrive CaptchaId = %s, CaptchaValue= %s", uar.CaptchaId, uar.CaptchaValue)
	if !cpt.Verify(uar.CaptchaId, uar.CaptchaValue) {
		ctx.JSON(http.StatusBadRequest, rest.INVALID_CAPTCHA)
		return
	}

	valid := validation.Validation{}
	valid.Email(uar.Email, "Email")
	valid.Match(uar.Username, rest.ValidPasswd,
		"Username").Message(rest.UsernamePrompt)
	valid.Match(uar.Passwd, rest.ValidPasswd,
		"Passwd").Message(rest.PasswdPrompt)
	if !validMember(ctx, &valid) {
		return
	}

	// check user whether existed
	u := &models.User{}
	if err := u.Find(uar.Email, uar.Username, ""); err != orm.ErrNoRows {
		ctx.JSON(http.StatusBadRequest, rest.INVALID_SIGNUP)
		return
	}

	// check reserve users
	if _, ok := rest.ReserveUsers[uar.Username]; ok {
		ctx.JSON(http.StatusBadRequest, rest.INVALID_SIGNUP)
		return
	}

	// generate password mask
	pwd, salt := tkits.GenPasswd(uar.Passwd, 8)
	u.Salt = salt
	u.Password = pwd
	u.Updated = time.Now()
	u.Username = uar.Username
	u.Email = uar.Email
	if id, err := u.Insert(); err != nil {
		ctx.JSON(http.StatusInternalServerError, tkits.DB_ERROR)
		return
	} else {
		u.Id = id
	}

	// generate a token
	if token, err := as.GenUserToken(ctx.RemoteAddr(), u.Id, 15, rest.TokenUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, tkits.SYS_ERROR)
		return
	} else {
		rsp := &rest.UserAddRsp{u.Id, u.Username, token}

		// set some cookies
		if uar.CookieMaxAge == 0 {
			uar.CookieMaxAge = 60 * 60 * 12 //half of one day
		}

		suid := fmt.Sprintf("%v", u.Id)
		ctx.SetCookie("token", token, uar.CookieMaxAge)
		ctx.SetCookie("uid", suid, uar.CookieMaxAge)

		ctx.JSON(http.StatusOK, rsp)
	}
}
