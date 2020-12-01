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
	var cijenkins_offline []models.CiJenkins
	var cijenkins_update []models.CiJenkins
	var cijenkins_security []models.CiJenkins
	o.QueryTable("ci_jenkins").All(&cijenkins, "Name", "Domain", "Location", "Status","Environ","Type", "Currver", "Url")
	o.QueryTable("ci_jenkins").Filter("status", "Offline").All(&cijenkins_offline, "Name", "Domain", "Location", "Type", "Currver", "Url")
	o.QueryTable("ci_jenkins").Filter("isupdated", "N").All(&cijenkins_update, "Name", "Domain", "Location", "Type", "Currver", "Url")
	o.QueryTable("ci_jenkins").Filter("issecure", "N").All(&cijenkins_security, "Name", "Domain", "Location", "Type", "Currver", "Url")	

	c.Data["s"] = cijenkins
	c.Data["total"] = len(cijenkins)
	c.Data["offline"] = len(cijenkins_offline)
	c.Data["update"] = len(cijenkins_update)
	c.Data["security"] = len(cijenkins_security)	
	c.TplName = "index.html"
}
