package models

import (
	"fmt"

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

func (t *TestModel) GetList(page int, size int) (count int64, list []*TestModel) {
	o := orm.NewOrm()
	o.QueryTable(t.TableName()).Limit(size, (page-1)*size).All(&list)
	count, err := o.QueryTable(t.TableName()).Count()
	if err != nil {
		fmt.Println(err)
	}
	return

}
