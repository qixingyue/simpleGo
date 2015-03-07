package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["leafApi/controllers:AuthvdiskController"] = append(beego.GlobalControllerRouter["leafApi/controllers:AuthvdiskController"],
		beego.ControllerComments{
			"AddPost",
			`/authvdisk/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["leafApi/controllers:TaskController"] = append(beego.GlobalControllerRouter["leafApi/controllers:TaskController"],
		beego.ControllerComments{
			"AddTask",
			`/task/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["leafApi/controllers:QueueController"] = append(beego.GlobalControllerRouter["leafApi/controllers:QueueController"],
		beego.ControllerComments{
			"QueueList",
			`/queue/list`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["leafApi/controllers:QueueController"] = append(beego.GlobalControllerRouter["leafApi/controllers:QueueController"],
		beego.ControllerComments{
			"Delqueue",
			`/queue/remove`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["leafApi/controllers:QueueController"] = append(beego.GlobalControllerRouter["leafApi/controllers:QueueController"],
		beego.ControllerComments{
			"Status",
			`/queue/status`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["leafApi/controllers:QueueController"] = append(beego.GlobalControllerRouter["leafApi/controllers:QueueController"],
		beego.ControllerComments{
			"CreateQueue",
			`/queue/create`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["leafApi/controllers:DownloadController"] = append(beego.GlobalControllerRouter["leafApi/controllers:DownloadController"],
		beego.ControllerComments{
			"AddPost",
			`/download/add`,
			[]string{"post"},
			nil})

}
