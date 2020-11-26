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
	//var ids []int
	//var names []string
	//orm.RegisterModel(new(CiJenkins))
	o := orm.NewOrm()

	// Only return Id and Title
	var cijenkins []models.CiJenkins
	o.QueryTable("ci_jenkins").All(&cijenkins, "Name", "Domain", "Location", "Type", "Currver", "Url")

	//var query = fmt.Sprintf("SELECT * FROM public.test")
	//var num, err = o.Raw(query).QueryRows(&ids, &names)

	//if err == nil {
	//fmt.Println("postgres row affected nums: ", num)
	//fmt.Println(ids, names)
	//}

	//	ss := []string{"a", "b", "c"}
	c.Data["s"] = cijenkins
	c.TplName = "index.html"
}
