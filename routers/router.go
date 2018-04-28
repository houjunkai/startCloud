package routers

import (
	"github.com/astaxie/beego"
	"startCloud/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.UserPermission{}, "get:Login")
}
