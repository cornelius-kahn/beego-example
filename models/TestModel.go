package models

import (
	"github.com/astaxie/beego/orm"
)

type TestModel struct {
	Id   int    `json:"id"`
	Info string `json:"info"`
	Time int    `json:"time"`
}

func init() {
	orm.RegisterModel(new(TestModel))
}

func NewTestModel() *TestModel {
	return &TestModel{}
}

func (t *TestModel) TableName() string {
	return "info"
}

func (t *TestModel) GetList() (list []*TestModel) {
	o := orm.NewOrm()
	o.QueryTable(t.TableName()).All(&list)
	return
}
