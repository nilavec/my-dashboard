package controllers

import (
	"ah-dashboard/models"
	"strings"	
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CiJenkinsFull struct {
	Id       int
	Name     string
	Domain   string
	Location string
	Type     string
	Currver  string
	Url      string
	Cveids   []string
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	
	o := orm.NewOrm()

	var cijenkins []models.CiJenkins
	var cijenkins_offline []models.CiJenkins
	var cijenkins_update []models.CiJenkins
	var cijenkins_security []models.CiJenkins
	o.QueryTable("ci_jenkins").All(&cijenkins, "Name", "Domain", "Location", "Status","Environ","Type", "Currver", "Url", "Cluster", "Nspace", "Descr")
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
func (c *MainController) GetAll() {	
	o := orm.NewOrm()
	var cijenkins []models.CiJenkins
	o.QueryTable("ci_jenkins").All(&cijenkins, "Name", "Domain", "Location", "Status","Environ","Type", "Currver", "Url", "Cluster", "Nspace", "Descr")	
	c.Data["s"] = cijenkins
	c.TplName = "all.html"
}
func (c *MainController) GetOffline() {	
	o := orm.NewOrm()
	var cijenkins_offline []models.CiJenkins
	o.QueryTable("ci_jenkins").Filter("status", "Offline").All(&cijenkins_offline, "Name", "Domain", "Location", "Type", "Currver", "Url", "Cluster", "Nspace", "Descr")	
	c.Data["s"] = cijenkins_offline
	c.TplName = "offline.html"
}
func (c *MainController) GetUpdate() {	
	o := orm.NewOrm()
	var cijenkins_update []models.CiJenkins
	o.QueryTable("ci_jenkins").Filter("isupdated", "N").All(&cijenkins_update, "Name", "Domain", "Location", "Type", "Currver", "Lts", "Url", "Cluster", "Nspace", "Descr")
	c.Data["s"] = cijenkins_update	
	c.TplName = "update.html"
}
func (c *MainController) GetSecurity() {	
	o := orm.NewOrm()
	var cijenkins_security []models.CiJenkins
	var cijenkins_security_full []CiJenkinsFull	
	o.QueryTable("ci_jenkins").Filter("issecure", "N").All(&cijenkins_security, "Name", "Domain", "Location", "Type", "Currver", "Url", "Cveids")
	for i, s := range cijenkins_security {
		temp := strings.Split(s.Cveids, "\n")
		mystruct := CiJenkinsFull{Name: cijenkins_security[i].Name, Domain: cijenkins_security[i].Domain, Location: cijenkins_security[i].Location, Type: cijenkins_security[i].Type, Currver: cijenkins_security[i].Currver, Url: cijenkins_security[i].Url, Cveids: temp}
		cijenkins_security_full = append(cijenkins_security_full, mystruct)
	}	
	c.Data["s"] = cijenkins_security_full
	c.TplName = "security.html"
}
