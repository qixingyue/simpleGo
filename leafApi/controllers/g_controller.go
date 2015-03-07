package controllers

import (
	"github.com/astaxie/beego"
)

type G_Controller struct {
	beego.Controller
}

type ResultInfo struct {
	ErrorString string
	Result      string
	Addition    string
}

func ErrorJsonEnd(c *G_Controller, message string) {
	c.Data["json"] = createErrorObj(message)
	c.ServeJson()
}

func OKJsonEnd(c *G_Controller) {
	c.Data["json"] = createOKObj()
	c.ServeJson()
}

func OKJsonData(c *G_Controller, d interface{}) {
	c.Data["json"] = d
	c.ServeJson()
}

func OKJsonEndAddition(c *G_Controller, additionString string) {
	c.Data["json"] = createOKObjAddition(additionString)
	c.ServeJson()
}

func createErrorObj(message string) *ResultInfo {
	rif := new(ResultInfo)
	rif.ErrorString = message
	rif.Result = "failed"
	rif.Addition = ""
	return rif
}

func createOKObj() *ResultInfo {
	rif := new(ResultInfo)
	rif.ErrorString = ""
	rif.Result = "ok"
	rif.Addition = ""
	return rif
}

func createOKObjAddition(additionString string) *ResultInfo {
	rif := new(ResultInfo)
	rif.ErrorString = ""
	rif.Result = "ok"
	rif.Addition = additionString
	return rif
}
