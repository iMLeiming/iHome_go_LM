package routers

import (
	"github.com/astaxie/beego"
	"iHome_go_LM/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
