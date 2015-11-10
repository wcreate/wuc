package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/wcreate/tkits"
	"github.com/wcreate/wuc/models"
	"github.com/wcreate/wuc/rest"
	"gopkg.in/macaron.v1"

	log "github.com/Sirupsen/logrus"
)

var (
	DefaultAvatar = "/static/imgs/davatar.jpg"
)

// POST /api/user/avatar/:uid
func UploadAvatar(ctx *macaron.Context, as rest.AuthService, ut *rest.UserToken) {
	// retrive uid and check auth
	uid, ok := getUidWithAuth(ctx, as, ut, rest.DummyOptId)
	if !ok {
		return
	}

	// 2.0
	ui := &models.UserInfo{User: &models.User{Id: uid}}
	if err := ui.Read("User"); err == orm.ErrNoRows {
		ctx.JSON(http.StatusNotFound, rest.INVALID_USER)
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, tkits.DB_ERROR)
		return
	}

	// 3.0
	f, h, err := ctx.GetFile("avatar")
	log.Debugln("filename=", h.Filename)
	defer f.Close()

	// FIXME: the workdir should app root path
	ext := h.Filename[strings.LastIndex(h.Filename, ".")+1:]
	path := "upload/avatar/" + dateFormat(time.Now(), "y/m/d/h/")
	os.MkdirAll(path, 0744)
	tofile := path + strconv.Itoa(int(uid)) + "." + ext

	tf, err := os.OpenFile(tofile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Errorln("write failed, ", err.Error())
		return
	}
	defer tf.Close()
	io.Copy(tf, f)

	// remove the
	if ui.Avatar != "" && ui.Avatar != DefaultAvatar {
		os.Remove(ui.Avatar[1:])
	}
	ui.Avatar = "/" + tofile
	if _, err := ui.Update("Avatar"); err != nil {
		ctx.JSON(http.StatusInternalServerError, tkits.DB_ERROR)
		return
	}
	rsp := &rest.UploadAvatorRsp{ImgUrl: ui.Avatar}
	ctx.JSON(http.StatusOK, rsp)
}

// Date takes a PHP like date func to Go's time format.
func dateFormat(t time.Time, format string) string {
	res := strings.Replace(format, "MM", t.Format("01"), -1)
	res = strings.Replace(res, "M", t.Format("1"), -1)
	res = strings.Replace(res, "DD", t.Format("02"), -1)
	res = strings.Replace(res, "D", t.Format("2"), -1)
	res = strings.Replace(res, "YYYY", t.Format("2006"), -1)
	res = strings.Replace(res, "YY", t.Format("06"), -1)
	res = strings.Replace(res, "HH", fmt.Sprintf("%02d", t.Hour()), -1)
	res = strings.Replace(res, "H", fmt.Sprintf("%d", t.Hour()), -1)
	res = strings.Replace(res, "hh", t.Format("03"), -1)
	res = strings.Replace(res, "h", t.Format("3"), -1)
	res = strings.Replace(res, "mm", t.Format("04"), -1)
	res = strings.Replace(res, "m", t.Format("4"), -1)
	res = strings.Replace(res, "ss", t.Format("05"), -1)
	res = strings.Replace(res, "s", t.Format("5"), -1)
	return res
}
