package controllers

import (
	"beego-example/commons/responses"
	"beego-example/models"

	"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

func (c *MainController) Test() {
	page, errPage := c.GetInt("page")
	size, errSize := c.GetInt("size")
	if errPage != nil || errSize != nil {
		fmt.Println(errPage)
		fmt.Println(errSize)
	}
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 5
	}

	fmt.Println(page)
	fmt.Println(size)
	count, list := models.NewTestModel().GetList(page, size)

	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}

	c.Data["json"] = responses.SuccessDataPageReturn(count, list)
	c.ServeJSON()
}
