package routers

import (
	"OrderSystem/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.OrderController{})
}
