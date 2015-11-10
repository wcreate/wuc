package handler

import (
	"net/http"

	"github.com/wcreate/wuc/rest"
	"gopkg.in/macaron.v1"
)

// POST /api/user/logout?uid
func LogoutUser(ctx *macaron.Context, as rest.AuthService, ut *rest.UserToken) {
	ctx.Status(http.StatusOK)
}
