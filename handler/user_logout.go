package handler

import (
	"net/http"

	"github.com/wcreate/tkits"
	"gopkg.in/macaron.v1"
)

// POST /api/user/logout?uid
func LogoutUser(ctx *macaron.Context, as tkits.AuthService, ut *tkits.UserToken) {
	ctx.Status(http.StatusOK)
}
