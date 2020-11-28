package routers

import (
	"ah-dashboard/controllers"
	"ah-dashboard/db"

	"github.com/astaxie/beego"
)

func init() {
	db.InitDb()
	BasePath := beego.AppConfig.String("basePath")

	ns :=
		beego.NewNamespace(fmt.Sprintf("/%s", BasePath),
			beego.NSRouter("/", &controllers.MainController{}),
		)
	beego.AddNamespace(ns)	
	//beego.Router("/toolzilla", &controllers.MainController{})
}
