package controllers

import (
	"ah-dashboard/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	
	o := orm.NewOrm()

	var cijenkins []models.CiJenkins
	o.QueryTable("ci_jenkins").All(&cijenkins, "Name", "Domain", "Location", "Status","Environ","Type", "Currver", "Url")

	c.Data["s"] = cijenkins
	c.TplName = "index.html"
}
