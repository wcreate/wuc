package handler

import (
	"encoding/json"
	"net/http"

	"github.com/wcreate/tkits"

	"gopkg.in/macaron.v1"
)

func checkAuth(ctx *macaron.Context, uid int64) bool {
	return tkits.CheckAuth(ctx, uid)
}

func getClientIP(ctx *macaron.Context) string {
	return ctx.RemoteAddr()
}

func getUidWithAuth(ctx *macaron.Context) (int64, bool) {
	uid := ctx.ParamsInt64("uid")
	if uid == 0 {
		ctx.JSON(http.StatusBadRequest, tkits.INVALID_URL)
		return -1, false
	}

	// 1.0
	if !checkAuth(ctx, uid) {
		return -1, false
	}
	return uid, true
}

func getUidAndBodyWithAuth(ctx *macaron.Context, v interface{}) (int64, bool) {
	uid, ok := getUidWithAuth(ctx)
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
		ctx.JSON(http.StatusBadRequest, tkits.INVALID_BODY)
		return false
	}

	if err := json.Unmarshal(body, v); err == nil {
		ctx.JSON(http.StatusBadRequest, tkits.INVALID_BODY)
		return false
	}

	return true
}
