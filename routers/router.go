package routers

import (
	"curd/controllers"
	"github.com/astaxie/beego"
)

func init() {

    userController := &controllers.UserConteroller{}

	beego.Router("/", userController, "get:List")
    beego.Router("/admin/user/save",userController, "post:Save")
	beego.Router("/admin/user/delete", userController, "post:Delete")
	beego.Router("/admin/user/edit", userController, "get:Edit")

}
