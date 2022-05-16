package routers

import (
	"OrderSystem/controllers"

	//"github.com/astaxie/beego"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//beego.Include(&controllers.OrderController{})
	web.Router("/cache/get", &controllers.OrderController{}, "get:Get")
	web.Router("/cache/set", &controllers.OrderController{}, "post:Set")

}
