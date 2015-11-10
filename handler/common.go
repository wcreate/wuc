package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego/validation"
	"github.com/wcreate/tkits"
	"github.com/wcreate/wuc/rest"

	"gopkg.in/macaron.v1"
)

func getUidWithAuth(ctx *macaron.Context, as rest.AuthService, ut *rest.UserToken, opId rest.OperationId) (int64, bool) {
	uid := ctx.ParamsInt64("uid")
	if uid == 0 {
		ctx.JSON(http.StatusBadRequest, tkits.INVALID_URL)
		return -1, false
	}

	// 1.0
	if !as.Authorize(ut, uid, opId) {
		return uid, false
	}
	return uid, true
}

func getUidAndBodyWithAuth(ctx *macaron.Context, as rest.AuthService, ut *rest.UserToken, opId rest.OperationId, v interface{}) (int64, bool) {
	uid, ok := getUidWithAuth(ctx, as, ut, opId)
	if !ok {
		return uid, ok
	}

	ok = getBody(ctx, v)
	if !ok {
		return uid, ok
	}

	return uid, true
}

func getBody(ctx *macaron.Context, v interface{}) bool {
	body, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, tkits.INVALID_BODY)
		return false
	}

	log.Debugf("retrive body = %s", string(body))
	if err := json.Unmarshal(body, v); err != nil {
		log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, tkits.INVALID_BODY)
		return false
	}

	return true
}

func validReq(ctx *macaron.Context, v interface{}) bool {
	valid := validation.Validation{}
	ok, _ := valid.Valid(&v)
	if !ok {
		detail := ""
		for _, err := range valid.Errors {
			detail += fmt.Sprintf("[key=%s, message=%s, value=%v]", err.Key, err.Message, err.Value)
		}
		ctx.JSON(http.StatusBadRequest, tkits.Error{
			ErrorMsg: "invalid_request",
			Detail:   detail,
		})
		return false
	}
	return true
}

func validMember(ctx *macaron.Context, valid *validation.Validation) bool {
	if valid.HasErrors() {
		detail := ""
		for _, err := range valid.Errors {
			detail += fmt.Sprintf("[key=%s, message=%s, value=%v]", err.Key, err.Message, err.Value)
		}
		ctx.JSON(http.StatusBadRequest, tkits.Error{
			ErrorMsg: "invalid_request",
			Detail:   detail,
		})
		return false
	}
	return true
}
