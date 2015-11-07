package api

import "github.com/wcreate/wuc/models"

var (
	INVALID_URL = Error{
		"invalid url!",
		"url is invalid or method is not correct.",
	}

	INVALID_AUTH = Error{
		"invalid request!",
		"not found Authorization in header or the value is invalid.",
	}

	INVALID_BODY = Error{
		"invalid request!",
		"request body is not correct for this url.",
	}

	DB_ERROR = Error{
		"system error!",
		"operate db failed.",
	}

	SYS_ERROR = Error{
		"system error!",
		"unkown error.",
	}

	INVALID_USER = Error{
		"invalid request!",
		"invalid user or password.",
	}

	INVALID_CAPTCHA = Error{
		"invalid captcha!",
		"invalid captcha.",
	}
)

// Common Error Response
type Error struct {
	ErrorMsg string `json:"error"`
	Detail   string `json:"detail"`
}

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
