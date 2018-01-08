package routers

import (
	"iHome_go_LM/iHome_go_LM/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
