package controllers

import (
	"github.com/astaxie/beego"
	v "github.com/astaxie/beego/validation"
	m "leafApi/models"
)

// oprations for Download
type DownloadController struct {
	beego.Controller
}

func (c *DownloadController) URLMapping() {
	c.Mapping("Post", c.Post)
}

type DownloadResultInfo struct {
	ErrorString string
	RsyncPath   string
	Result      string
}

// @Title Post
// @Description create Download
// @Param	body		body 	models.Download	true		"body for Download content"
// @Success 200 {int} models.Download.Id
// @Failure 403 body is empty
// @router /download/add [post]
func (c *DownloadController) AddPost() {
	valid := v.Validation{}
	url, uniqueId, aimMd5 := c.GetString("url"), c.GetString("uniqueId"), c.GetString("aimMd5")
	valid.Required(url, "url")
	valid.Required(uniqueId, "uniqueId")
	valid.Required(aimMd5, "aimMd5")
	if valid.HasErrors() {
		dri := new(DownloadResultInfo)
		dri.ErrorString = "data error"
		c.Data["json"] = dri
		c.ServeJson()
		return
	}
	dri := new(DownloadResultInfo)
	dm := new(m.DownloadInfo)
	dm.Url = url
	dm.Uniqueid = uniqueId
	dm.AimMd5 = aimMd5
	dri.RsyncPath, dri.ErrorString = dm.WriteRedis()
	dri.Result = "ok"
	c.Data["json"] = &dri
	c.ServeJson()
}
