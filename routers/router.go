package routers

import (
	"github.com/astaxie/beego"
	"github.com/grayzone/three/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/loaddata", &controllers.MainController{}, "POST:LoadData")
}
