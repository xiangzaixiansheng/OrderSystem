package main

import (
	_ "OrderSystem/routers"
	"fmt"

	"github.com/beego/beego/logs"
	"github.com/beego/beego/orm"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/go-sql-driver/mysql"
)

var (
	ConfigFile = "./conf/app.conf"
)

func main() {
	err := web.LoadAppConfig("ini", ConfigFile)
	if err != nil {
		logs.Critical("An error occurred:", err)
		panic(err)
	}
	defaultdb, _ := web.AppConfig.String("defaultdb")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	fmt.Println(defaultdb)
	orm.RegisterDataBase("default", "mysql", defaultdb, 30, 30)

	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	if web.BConfig.RunMode == "dev" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	web.Run()
}
