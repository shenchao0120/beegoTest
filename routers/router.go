package routers

import (
	"chaoshen.com/beegoTest/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
