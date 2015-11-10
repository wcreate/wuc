package handler

import (
	"github.com/wcreate/tkits"
	"github.com/wcreate/wuc/models"
	"github.com/wcreate/wuc/rest"
	"gopkg.in/macaron.v1"

	"net/http"
)

// DELETE /api/user/:uid/
func DeleteUser(ctx *macaron.Context, as rest.AuthService, ut *rest.UserToken) {
	// 1.0
	uid, ok := getUidWithAuth(ctx, as, ut, rest.DummyOptId)
	if !ok {
		return
	}

	u := &models.User{Id: uid}
	if err := u.DeleteAll(); err != nil {
		ctx.JSON(http.StatusInternalServerError, tkits.DB_ERROR)
		return
	}

	ctx.Status(http.StatusOK)
}
