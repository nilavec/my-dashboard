package db

import (
	"ah-dashboard/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

//var log = logger.GetLogger()
var ids []int
var names []string

func InitDb() {
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	fmt.Println("Postgres driver registered.")
	checkErr(err)

	pgUser := beego.AppConfig.String("pgUser")
	pgPassword := beego.AppConfig.String("pgPassword")
	pgHost := beego.AppConfig.String("pgHost")
	pgDatabase := beego.AppConfig.String("pgDatabase")
	pgPort := beego.AppConfig.String("pgPort")
	//pgSchema := beego.AppConfig.String("pgSchema")

	fmt.Println(pgUser, pgPassword, pgDatabase, pgHost, pgPort)

	params := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		pgUser, pgPassword, pgHost, pgPort, pgDatabase)

	err = orm.RegisterDataBase("default", "postgres", params)
	checkErr(err)
	fmt.Println("Connection to database is established.",
		zap.String("host", pgHost),
		zap.String("port", pgPort))

	orm.RegisterModel(new(models.CiJenkins))

	o := orm.NewOrm()
	var num int64

	var query = fmt.Sprintf("SELECT * FROM public.test")
	num, err = o.Raw(query).QueryRows(&ids, &names)

	if err == nil {
		fmt.Println("postgres row affected nums: ", num)
		fmt.Println(ids, names)
	}

	//db, err := orm.GetDB("default")
	//checkErr(err)
	//driver, err := postgres.WithInstance(db, &postgres.Config{})
}

func checkErr(err error) {
	if err != nil {
		handleErr(err)
	}
}

func handleErr(err error) {
	if err.Error() == "no change" {
		fmt.Println("Warning from db migration", zap.Error(err))
	} else {
		fmt.Println("An error has occurred during migration", zap.Error(err))
	}
}

func getNames() []string {
	o := orm.NewOrm()

	var query = fmt.Sprintf("SELECT * FROM public.test")
	var num, err = o.Raw(query).QueryRows(&ids, &names)

	if err == nil {
		fmt.Println("postgres row affected nums: ", num)
		fmt.Println(ids, names)
	}
	return names

}
