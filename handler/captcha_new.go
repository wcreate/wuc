package handler

import (
	"fmt"
	"net/http"

	"github.com/go-macaron/cache"
	"github.com/go-macaron/captcha"
	"github.com/wcreate/wuc/api"
	"gopkg.in/macaron.v1"
)

// GET /captcha/new
func GetCaptcha(ctx *macaron.Context, cpt *captcha.Captcha, cache cache.Cache) {

	// only allow 10 times request in one second
	cip := ctx.RemoteAddr()
	times, _ := cache.Get(cip).(int)
	if times > 10 {
		ctx.Status(http.StatusForbidden)
		return
	}
	cache.Put(cip, times+1, 1)

	// create the captcha
	v, err := cpt.CreateCaptcha()
	if err != nil {
		panic(fmt.Errorf("fail to create captcha: %v", err))
	}

	rsp := &api.CaptchaRsp{
		cpt.FieldIdName,
		v,
		fmt.Sprintf("%s%s%s.png", cpt.SubURL, cpt.URLPrefix, v),
	}

	ctx.JSON(http.StatusOK, rsp)
}
