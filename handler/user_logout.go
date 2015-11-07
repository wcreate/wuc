package handler

import (
	"net/http"

	"gopkg.in/macaron.v1"
)

// POST /api/user/logout?uid
func LogoutUser(ctx *macaron.Context) {
	ctx.Status(http.StatusOK)
}
