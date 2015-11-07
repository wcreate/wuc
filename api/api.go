package api

import (
	"github.com/wcreate/tkits"
	"github.com/wcreate/wuc/models"
)

var (
	INVALID_USER = tkits.Error{
		"invalid request!",
		"invalid user or password.",
	}

	INVALID_CAPTCHA = tkits.Error{
		"invalid captcha!",
		"invalid captcha.",
	}
)

// Modify Password Request
type ModifyPasswordReq struct {
	Uid       string `json:"uid"`
	OldPasswd string `json:"old_password"`
	NewPasswd string `json:"new_password"`
}

// Modify UserInfo Request
type ModifyUserInfoReq struct {
	Uid int64 `json:"uid"`
	models.UserInfo
}

// User login Request
type UserLoginReq struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	Mobile      string `json:"mobile"`
	Passwd      string `json:"password"`
	CaptchaId   string `json:"captcha_id"`
	CaptchaName string `json:"captcha_name"`
}

// User login Response
type UserLoginRsp struct {
	Uid      int64  `json:"uid"`
	Username string `json:"username"`
	Token    string `json:"token"`
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
