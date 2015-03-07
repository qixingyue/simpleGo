package main

import (
	"github.com/astaxie/beego"
	_ "leafApi/docs"
	_ "leafApi/routers"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.SetStaticPath("/public", "public")
	beego.Run()
}
