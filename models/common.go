package models

import "github.com/astaxie/beego/orm"

type Transaction struct {
	o   orm.Ormer
	err error
}

func (t *Transaction) Begin() {
	t.o = orm.NewOrm()
	t.o.Begin()
}

func (t *Transaction) PostDefer() {
	if t.err == nil {
		t.o.Commit()
	} else {
		t.o.Rollback()
	}
}
