package routers

import (
	"github.com/astaxie/beego"
	"sdbm/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
