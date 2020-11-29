package models

type CiJenkins struct {
	Id       int    `json:"id" orm:"column(id)"`
	Name     string `json:"name" orm:"column(name)"`
	Domain   string `json:"domain" orm:"column(domain)"`
	Location string `json:"location" orm:"column(location)"`
	Status string `json:"location" orm:"column(status)"`
	Environ string `json:"location" orm:"column(environ)"`	
	Type     string `json:"type" orm:"column(type)"`
	Currver  string `json:"currver" orm:"column(currver)"`
	Url      string `json:"url" orm:"column(url)"`
}

func (u *CiJenkins) TableName() string {
	// db table name
	return "ci_jenkins"
}
