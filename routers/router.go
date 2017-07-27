package routers

import (
	"github.com/astaxie/beego"
	"github.com/grayzone/three/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
