package controllers

import (
	"github.com/astaxie/beego"
	v "github.com/astaxie/beego/validation"
	m "leafApi/models"
)

// oprations for Authvdisk
type AuthvdiskController struct {
	beego.Controller
}

type AuthvdiskResultInfo struct {
	ErrorString string
	//RsyncPath   string
	Result string
}

// @Title Post
// @Description create AuthVdisk
// @Success 200 {int} models.Download.Id
// @Failure 403 body is empty
// @router /authvdisk/add [post]
func (c *AuthvdiskController) AddPost() {
	valid := v.Validation{}
	jsondata, uniqueId := c.GetString("jsondata"), c.GetString("uniqueId")
	valid.Required(jsondata, "jsondata")
	valid.Required(uniqueId, "uniqueId")
	if valid.HasErrors() {
		dri := new(AuthvdiskResultInfo)
		dri.ErrorString = "data error"
		c.Data["json"] = dri
		c.ServeJson()
		return
	}

	dri := new(AuthvdiskResultInfo)

	dm := new(m.AuthvdiskM)
	dm.Uniqueid = uniqueId
	dm.Jsondata = jsondata

	dri.Result, dri.ErrorString = dm.WriteRedis()
	c.Data["json"] = &dri
	c.ServeJson()
}
