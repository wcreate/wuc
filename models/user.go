package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// 用户基本信息
type User struct {
	Id       int64  `json:"uid" orm:"pk;auto"`
	Pid      int64  `json:"pid"`                      // 用在归属地 归属学校 归属组织 等方面
	Email    string `json:"email" orm:"index;unique"` // 电邮
	Salt     string `json:"-"`                        // 加密的盐值``
	Password string `json:"-"`                        // 加密之后的密文
	Cfmcode  string `json:"-"`                        // 邮件确认码

	Username string `json:"username" orm:"index;unique"` // 用户名
	Mobile   string `json:"mobile" orm:"index"`          // 手机

	Ctype int64 `json:"ctype" orm:"index"` // 类型
	Role  int64 `json:"role" orm:"index"`  // 角色

	Created time.Time `json:"created" orm:"auto_now_add;type(datetime)"` // 创建时间
	Updated time.Time `json:"updated" orm:"type(datetime)"`              // 资料更新时间

	LastLoginTime time.Time `json:"login_time" orm:"type(datetime);null"` // 最后登录时间
	LastLoginIp   string    `json:"login_ip" valid:"IP"`                  // 最后登录的IP
	LoginCount    int64     `json:"login_times"`                          // 登录次数

	UserInfo *UserInfo `json:"info" orm:"null;reverse(one)"` // 设置反向关系(可选)
}

// 用户的详细信息
type UserInfo struct {
	Id   int64 `json:"-" orm:"pk;auto"`
	User *User `json:"-" orm:"null;rel(one)"` // OneToOne relation

	Realname string `json:"real_name" valid:"MaxSize(100)"`                  // 真实姓名
	Content  string `json:"signsure" orm:"size(1024)" valid:"MaxSize(1024)"` // 个人签名
	Avatar   string `json:"avatar"`                                          // 头像地址 48*48

	Gender int       `json:"gender" valid:"Range(0,2)"`           // 0：Unknown, 1: Male， 2：Female
	Birth  time.Time `json:"birth" orm:"auto_now_add;type(date)"` // 生日

	Province string `json:"province" valid:"MaxSize(32)"` // 省份
	City     string `json:"city" valid:"MaxSize(100)"`    // 城市
	Company  string `json:"company" valid:"MaxSize(255)"` // 公司
	Address  string `json:"address" valid:"MaxSize(255)"` // 地址
	Website  string `json:"website" valid:"MaxSize(255)"` // 个人官网

}

//-----------------------------------------------------------------------------
func (u *User) Read(fields ...string) error {
	err := orm.NewOrm().Read(u, fields...)
	return err
}

func (u *User) Insert() error {
	_, err := orm.NewOrm().Insert(u)
	return err
}

func (u *User) Update(fields ...string) (int64, error) {
	row, err := orm.NewOrm().Update(u, fields...)
	return row, err
}

func (u *User) Delete() error {
	_, err := orm.NewOrm().Delete(u)
	return err
}

func Users() orm.QuerySeter {
	return orm.NewOrm().QueryTable("User")
}

func (u *User) ReadOneOnly(fields ...string) error {
	fields = append(fields, "Id")
	return Users().Filter("Id", u.Id).One(u, fields...)
}

func (u *User) DeleteAll() error {
	t := &Transaction{}
	t.Begin()
	defer t.PostDefer()

	if _, t.err = UserInfos().Filter("Uid", u.Id).Delete(); t.err != nil {
		return t.err
	}

	if t.err = u.Delete(); t.err != nil {
		return t.err
	}

	return nil
}

func (u *User) Find(email, username, mobile string) error {
	if email != "" {
		u.Email = email
		return u.Read("Email")
	}
	if username != "" {
		u.Username = username
		return u.Read("Username")
	}
	if mobile != "" {
		u.Mobile = mobile
		return u.Read("Mobile")
	}
	return nil
}

func (u *User) QueryAll() error {
	return Users().Filter("Id", u.Id).RelatedSel("UserInfo").One(u)
}

//-----------------------------------------------------------------------------
func (u *UserInfo) Read(fields ...string) error {
	err := orm.NewOrm().Read(u, fields...)
	return err
}

func (u *UserInfo) Insert() error {
	_, err := orm.NewOrm().Insert(u)
	return err
}

func (u *UserInfo) Update(fields ...string) (int64, error) {
	row, err := orm.NewOrm().Update(u, fields...)
	return row, err
}

func (u *UserInfo) Delete() error {
	_, err := orm.NewOrm().Delete(u)
	return err
}

func UserInfos() orm.QuerySeter {
	return orm.NewOrm().QueryTable("UserInfo")
}

func (u *UserInfo) ReadOneOnly(fields ...string) error {
	fields = append(fields, "Id")
	return UserInfos().Filter("Id", u.Id).One(u, fields...)
}

func (ui *UserInfo) UpdateInfo() error {
	t := &Transaction{}
	t.Begin()
	defer t.PostDefer()

	if _, t.err = ui.Update(); t.err != nil {
		return t.err
	}

	u := &User{Id: ui.User.Id}
	u.Updated = time.Now()
	if _, t.err = u.Update(); t.err != nil {
		return t.err
	}

	return nil
}
