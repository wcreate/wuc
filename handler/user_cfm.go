package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/wcreate/tkits"
	"github.com/wcreate/wuc/models"
	"gopkg.in/macaron.v1"
)

var (
	CFM_MOD_SUBJET     = "登录邮箱地址修改"
	CFG_SIGNUP_SUBJECT = "登录邮箱地址确认"
)

var (
	emailbody = `<html><body><div style="margin: 0 auto; width: 580px; background: #FFF; box-shadow: 0 0 10px #333; text-align:left;">
        <div style="margin: 0 40px; color: #999; border-bottom: 1px dotted #DDD; padding: 40px 0 30px; font-size: 13px; text-align: center;">
            <a href="http://{{.domain}}"><img src="http://{{.domain}}/static/imgs/mail-logo.png" alt="{{.name}}"></a><br>
            {{.intro}}
        </div>
        <div style="padding: 30px 40px 40px;">%s 您好，您注册或修改了登录邮箱地址<br><br>请在 1 小时内点击此链接以完成修改
            <a style="color: #009A61; text-decoration: none;" href="http://{{.domain}}/user/cfm?t=mail&amp;uid={{.uid}}&amp;code={{.code}}">
              http://{{.domain}}/user/cfm?t=mail&amp;uid={{.uid}}&amp;code={{.code}}</a><br></div><div style="background: #EEE; border-top: 1px solid #DDD; text-align: center; height: 90px; line-height: 90px;">
            <a href="http://{{.domain}}/user/cfm?t=mail&amp;uid={{.uid}}&amp;code={{.code}}" style="padding: 8px 18px; background: #009A61; color: #FFF; text-decoration: none; border-radius: 3px;">完成确认 ➔</a>
        </div>
    </div></body></html>`
)

func SendCfmEmail(u *models.User, subject string) error {
	tmp := template.New("EmailBody")
	tmp, err := tmp.Parse(emailbody)
	if err != nil {
		return err
	}

	tmpmap := make(map[string]string)
	tmpmap["domain"] = tkits.WebDomain
	tmpmap["name"] = tkits.WebName
	tmpmap["intro"] = tkits.WebIntro
	tmpmap["uid"] = fmt.Sprintf("%v", u.Id)
	tmpmap["code"] = u.Cfmcode

	buf := bytes.NewBuffer(make([]byte, 1024))
	tmp.Execute(buf, tmpmap)
	body := buf.String()

	m := tkits.NewHTMLMessage(subject, body)
	m.From = tkits.EmailUser
	m.To = []string{u.Email}
	host := strings.Split(tkits.EmailHost, ":")[0]

	return tkits.SendEmail(tkits.EmailHost,
		smtp.PlainAuth("", tkits.EmailUser, tkits.EmailPasswd, host), m)
}

// Get /api/user/cfm?t=mail&uid=&code=
func ConfirmUser(ctx *macaron.Context) {
	// 1.0
	//uid := ctx.QueryInt64("uid")
	//code := ctx.QueryEscape("code")
	//u := &models.User{Id: uid}

	ctx.Status(http.StatusOK)
}
