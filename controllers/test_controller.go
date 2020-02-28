package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

type TestMap struct {
	Food   string
	Watch  string
	Listen string
}

type JSONS struct {
	Ret  string  `json:"ret"`
	Msg  string  `json:"msg"`
	Data TestMap `json:"data"`
}

func (c *TestController) Get() {
	data := &JSONS{"0", "ok", TestMap{"蛋糕", "电影", "音乐"}}
	c.Data["json"] = data
	c.ServeJSON()
}
