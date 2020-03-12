package main

import (
	_ "beego-example/commons/routers"

	"os"

	"github.com/astaxie/beego"
)

func main() {
	// can set this variable `GO_DEPLOY_ENV` in ~/.bash_profile or /etc/profile
	// such as `export GO_DEPLOY_ENV="prod"`
	// this variable has values: `prod`, `dev`, the default value is `dev`
	env := os.Getenv("GO_DEPLOY_ENV")

	// load config files
	beego.LoadAppConfig("ini", "conf/app.conf")
	if env == "prod" {
		beego.LoadAppConfig("ini", "conf/app-prod.conf")
	} else {
		beego.LoadAppConfig("ini", "conf/app-dev.conf")
	}

	beego.Run()
}
