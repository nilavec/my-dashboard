package routers

import (
	"ah-dashboard/controllers"
	"ah-dashboard/db"

	"github.com/astaxie/beego"
)

func init() {
	db.InitDb()
	beego.Router("/", &controllers.MainController{})
}
