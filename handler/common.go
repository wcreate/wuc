package handler

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/wcreate/wuc/api"
	"github.com/wcreate/wuc/security"
	"gopkg.in/macaron.v1"
)

func checkAuth(ctx *macaron.Context, uid int64) bool {
	auth := strings.TrimSpace(ctx.Header().Get("Authorization"))

	if !security.GetDefaultSimpleToken().Validate(auth, getClientIP(ctx), fmt.Sprintf("%v", uid)) {
		ctx.JSON(404, api.INVALID_AUTH)
		return false
	}
	return true
}

func getClientIP(ctx *macaron.Context) string {
	raddr := ctx.RemoteAddr()
	ripport := strings.SplitN(raddr, ":", 2)
	return ripport[0]
}

func getUidWithAuth(ctx *macaron.Context) (int64, bool) {
	uid := ctx.ParamsInt64("uid")
	if uid == 0 {
		ctx.JSON(400, api.INVALID_URL)
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
		ctx.JSON(400, api.INVALID_BODY)
		return false
	}

	if err := json.Unmarshal(body, v); err == nil {
		ctx.JSON(400, api.INVALID_BODY)
		return false
	}

	return true
}
