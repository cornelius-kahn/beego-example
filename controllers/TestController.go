package controllers

import (
	"beego-example/commons/responses"
	"beego-example/models"

	"fmt"

	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	test := models.NewTestModel().GetList()

	for i := 0; i < len(test); i++ {
		fmt.Print(test[i])
	}
	c.Data["json"] = responses.SuccessDataPageReturn(1, test)
	c.ServeJSON()
}
