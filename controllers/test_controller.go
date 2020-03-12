package controllers

import (
	"beego-example/commons/responses"

	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

type DataMap struct {
	Food   string
	Watch  string
	Listen string
}

func (c *TestController) Get() {
	result := DataMap{"蛋糕", "电影", "音乐"}
	data := responses.SuccessDataPageReturn(1, result)
	c.Data["json"] = data
	c.ServeJSON()
}
