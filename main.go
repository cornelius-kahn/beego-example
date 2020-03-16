package main

import (
	_ "beego-example/commons/routers"

	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
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
		orm.Debug = true
	}

	// register database driver
	orm.RegisterDriver(beego.AppConfig.String("db_adapter"), orm.DRMySQL)
	orm.RegisterDataBase(beego.AppConfig.String("db_conn_name"), beego.AppConfig.String("db_adapter"), beego.AppConfig.String("db_user")+":"+beego.AppConfig.String("db_password")+"@tcp("+beego.AppConfig.String("db_host")+":"+beego.AppConfig.String("db_port")+")/"+beego.AppConfig.String("db_database")+"?charset="+beego.AppConfig.String("db_charset"), beego.AppConfig.DefaultInt("db_max_idle", 10), beego.AppConfig.DefaultInt("db_max_conn", 50))

	// start server
	beego.Run()
}
