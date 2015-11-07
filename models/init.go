package models

import (
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wcreate/wuc/security"
	"gopkg.in/macaron.v1"
)

func init() {
	dbname := "default" // 数据库别名
	cfg := macaron.Config()
	web, err := cfg.GetSection("web")
	if err != nil {
		panic(err)
	}

	dbtype := web.Key("dbtype").String()
	log.Debugf("DB type is %s", dbtype)
	dbcfg, err := cfg.GetSection(dbtype)
	if err != nil {
		panic(err)
	}

	switch dbtype {
	case "mysql":
		var username string = dbcfg.Key("username").String()
		if username, err = security.GetDefaultCrypto().DecryptStr(username); err != nil {
			panic(err)
		}

		var password string = dbcfg.Key("password").String()
		if password, err = security.GetDefaultCrypto().DecryptStr(password); err != nil {
			panic(err)
		}

		url := dbcfg.Key("url").String()
		maxidle := dbcfg.Key("maxidle").MustInt(2)
		maxconn := dbcfg.Key("maxconn").MustInt(2)
		orm.RegisterDriver("mysql", orm.DR_MySQL)
		orm.RegisterDataBase(dbname, "mysql",
			username+":"+password+"@"+url,
			maxidle, maxconn)
	case "sqlite":
		url := dbcfg.Key("url").String()
		orm.RegisterDriver("sqlite3", orm.DR_Sqlite)
		orm.RegisterDataBase(dbname, "sqlite3", url)
	}

	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC
	orm.Debug = true
	orm.RegisterModel(new(User), new(UserInfo))

	force := false                      // drop table 后再建表
	sqllog := web.Key("sqlon").String() // 打印执行过程
	verbose := false
	if "on" == sqllog {
		verbose = true
	}
	// 遇到错误立即返回
	err = orm.RunSyncdb(dbname, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
