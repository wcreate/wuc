package handler

import (
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/wcreate/wuc/api"
	"github.com/wcreate/wuc/models"
	"gopkg.in/macaron.v1"
)

// Get /api/user/check?t=&v=
func CheckUser(ctx *macaron.Context) {
	// 1.0
	t := ctx.QueryEscape("t")
	v := ctx.QueryEscape("v")
	u := &models.User{}
	err := orm.ErrNoRows
	switch t {
	case "u":
		err = u.Find("", v, "")
	case "e":
		err = u.Find(v, "", "")
	case "m":
		err = u.Find("", "", v)
	default:
		break
	}
	if err != nil {
		ctx.JSON(http.StatusNotFound, api.INVALID_USER)
		return
	}
	ctx.Status(http.StatusOK)
}
