package routers

import (
	"beego-example/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/test", &controllers.MainController{}, "get:Test")
}
