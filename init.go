package wuc

import (
	"github.com/go-macaron/cache"
	"github.com/go-macaron/captcha"
	"github.com/wcreate/wuc/handler"
	"gopkg.in/macaron.v1"
)

func InitHandles(m *macaron.Macaron) {
	m.Use(cache.Cacher())
	m.Use(captcha.Captchaer(captcha.Options{
		URLPrefix:        "/captcha/img/", // 获取验证码图片的 URL 前缀，默认为 "/captcha/"
		FieldIdName:      "captcha_id",    // 表单隐藏元素的 ID 名称，默认为 "captcha_id"
		FieldCaptchaName: "captcha",       // 用户输入验证码值的元素 ID，默认为 "captcha"
		ChallengeNums:    6,               // 验证字符的个数，默认为 6
		Width:            240,             // 验证码图片的宽度，默认为 240 像素
		Height:           80,              // 验证码图片的高度，默认为 80 像素
		Expiration:       600,             // 验证码过期时间，默认为 600 秒
		CachePrefix:      "captcha_",      // 用于存储验证码正确值的 Cache 键名，默认为 "captcha_"
	}))
	m.Get("/captcha/new", handler.GetCaptcha)

	m.Delete("/api/user/:uid", handler.DeleteUser)

	m.Get("/api/user/info/:uid", handler.UserInfo)
	m.Put("/api/user/info/:uid", handler.ModifyUser)
	m.Post("/api/user/info/:uid", handler.ModifyUser)
	m.Put("/api/user/pwd/:uid", handler.ModifyPassword)

	m.Post("/api/user/login", handler.LoginUser)
}
