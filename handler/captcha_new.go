package handler

import (
	"fmt"

	"github.com/go-macaron/captcha"
	"github.com/wcreate/wuc/api"
	"gopkg.in/macaron.v1"
)

// GET /captcha/new
func GetCaptcha(ctx *macaron.Context, cpt *captcha.Captcha) {
	v, err := cpt.CreateCaptcha()
	if err != nil {
		panic(fmt.Errorf("fail to create captcha: %v", err))
	}

	rsp := &api.CaptchaRsp{
		cpt.FieldIdName,
		v,
		fmt.Sprintf("%s%s%s.png", cpt.SubURL, cpt.URLPrefix, v),
	}

	ctx.JSON(200, rsp)
}
