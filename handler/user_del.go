package handler

import (
	"github.com/wcreate/wuc/api"
	"github.com/wcreate/wuc/models"
	"gopkg.in/macaron.v1"
)

// DELETE /api/user/:uid/
func DeleteUser(ctx *macaron.Context) {
	// 1.0
	uid, ok := getUidWithAuth(ctx)
	if !ok {
		return
	}

	u := &models.User{Id: uid}
	if err := u.DeleteAll(); err != nil {
		ctx.JSON(500, api.DB_ERROR)
		return
	}
}
