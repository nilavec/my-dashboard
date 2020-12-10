package routers

import (
	"ah-dashboard/controllers"
	"ah-dashboard/db"
	"fmt"
	"github.com/astaxie/beego"
)

func init() {
	db.InitDb()
	BasePath := beego.AppConfig.String("basePath")
	beego.SetStaticPath(fmt.Sprintf("%s/static", BasePath), "static")
	ns :=
		beego.NewNamespace(fmt.Sprintf("/%s", BasePath),
			beego.NSRouter("/", &controllers.MainController{}),
			beego.NSRouter("/all", &controllers.MainController{}, "get:GetAll"),					   
			beego.NSRouter("/offline", &controllers.MainController{}, "get:GetOffline"),
			beego.NSRouter("/allgrafana", &controllers.MainController{}, "get:GetAllGrafana"),					   
			beego.NSRouter("/offlinegrafana", &controllers.MainController{}, "get:GetOfflineGrafana"),				   
			beego.NSRouter("/update", &controllers.MainController{}, "get:GetUpdate"),
			beego.NSRouter("/updategrafana", &controllers.MainController{}, "get:GetUpdateGrafana"),				   
			beego.NSRouter("/security", &controllers.MainController{}, "get:GetSecurity"),					   
		)
	beego.AddNamespace(ns)	
	//beego.Router("/toolzilla", &controllers.MainController{})
}
