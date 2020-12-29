package main

import (
	_ "curd/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/context"

)

func main() {

	orm.RegisterDataBase("default", "mysql", "root:iran@tcp(127.0.0.1:3307)/gocrud?charset=utf8")
	orm.Debug = true

	orm.RunSyncdb("default", false, true)

	var filterAdmin = func(ctx *context.Context) {
		url := ctx.Input.URL()
		logs.Info("##### filter url : %s", url)
		//TODO 如果判断用户未登录。

	}
	beego.InsertFilter("/admin/*", beego.BeforeExec, filterAdmin)

	beego.Run()
}

