package handler

import (
	"net/http"

	"gopkg.in/macaron.v1"
)

// Get /api/user/cfm?uid=&code=
func ConfirmUser(ctx *macaron.Context) {
	// 1.0
	//uid := ctx.QueryInt64("uid")
	//code := ctx.QueryEscape("code")
	//u := &models.User{Id: uid}

	ctx.Status(http.StatusOK)
}
