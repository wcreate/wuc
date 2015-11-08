package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/wcreate/tkits"
)

func init() {
	tkits.InitDB()
	orm.RegisterModel(new(User), new(UserInfo))
}
