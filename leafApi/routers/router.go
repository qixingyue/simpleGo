package routers

import (
	"github.com/astaxie/beego"
	"leafApi/controllers"
)

func init() {
	beego.Include(&controllers.DownloadController{})
	beego.Include(&controllers.AuthvdiskController{})
	beego.Include(&controllers.QueueController{})
	beego.Include(&controllers.TaskController{})
}
