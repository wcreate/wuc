package api

import (
	"regexp"

	"github.com/wcreate/tkits"
	"github.com/wcreate/wuc/models"
)

var (
	INVALID_USER = tkits.Error{
		"invalid_auth",
		"invalid user or password.",
	}

	INVALID_EMAIL = tkits.Error{
		"invalid_request",
		"invalid email.",
	}

	INVALID_SIGNUP = tkits.Error{
		"invalid_user",
		"email or user has existed.",
	}

	SEND_EMAIL_FAILED = tkits.Error{
		"sent_email_failed",
		"sent a comfirm email failed,please check whether the email is correct.",
	}

	INVALID_CAPTCHA = tkits.Error{
		"invalid_captcha",
		"the captcha id or value is invalid.",
	}
)

var (
	ValidUsername  = regexp.MustCompile("^[\\x{4e00}-\\x{9fa5}A-Z0-9a-z_-]{4,30}$")
	ValidPasswd    = regexp.MustCompile(`^[\@A-Za-z0-9\!\#\$\%\^\&\*\~\{\}\[\]\.\,\<\>\(\)\_\+\=]{4,30}$`)
	UsernamePrompt = "用户名是为永久性设定,不能少于4个字或多于30个字,请慎重考虑,不能为空!"
	PasswdPrompt   = "密码含有非法字符或密码过短(至少4~30位密码)!"

	ReserveUsers = map[string]string{
		"admin":         "admin",
		"administrator": "administrator",
		"home":          "home",
	}
)

// Modify Password Request
type ModifyPasswordReq struct {
	Uid       int64  `json:"uid"`
	OldPasswd string `json:"old_password"`
	NewPasswd string `json:"new_password"`
}

// Modify Email Request
type ModifyEmailReq struct {
	Uid      int64  `json:"uid"`
	OldEmail string `json:"old_email" valid:"Email; MaxSize(100)"`
	NewEmail string `json:"new_email" valid:"Email; MaxSize(100)"`
}

// Modify Email Request
type ModifyEmailRsp struct {
	Uid    int64  `json:"uid"`
	CfmUrl string `json:"cfm_url"`
}

//
type UserAddReq struct {
	Email        string `json:"email" valid:"Email; MaxSize(100)"`
	Username     string `json:"username" valid:"MaxSize(100)"`
	Passwd       string `json:"password"`
	CaptchaId    string `json:"captcha_id"`
	CaptchaValue string `json:"captcha_value"`
	CookieMaxAge int    `json:"cookie_maxage"`
}

type UserAddRsp struct {
	Uid      int64  `json:"uid"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

// Modify UserInfo Request
type ModifyUserInfoReq struct {
	Uid int64 `json:"uid"`
	models.UserInfo
}

// User login Request
type UserLoginReq struct {
	Email        string `json:"email"`
	Username     string `json:"username"`
	Mobile       string `json:"mobile"`
	Passwd       string `json:"password"`
	CaptchaId    string `json:"captcha_id"`
	CaptchaValue string `json:"captcha_value"`
	CookieMaxAge int    `json:"cookie_maxage"`
}

// User login Response
type UserLoginRsp struct {
	UserAddRsp
}

// User Info Response
type UserInfoRsp struct {
	models.User
}

type CaptchaRsp struct {
	FieldIdName string `json:"id_name"`
	Id          string `json:"id_value"`
	ImgUrl      string `json:"img_url"`
}
