package controllers

import (
	_ "fmt"
	v "github.com/astaxie/beego/validation"
	m "leafApi/models"
)

// oprations for QueueOperate
type QueueController struct {
	G_Controller
}

// @router /queue/list [get]
func (c *QueueController) QueueList() {
	qm := new(m.QueueModel)
	qlist := qm.GetQueues()
	OKJsonData(&c.G_Controller, qlist)
}

// @router /queue/remove [post]
func (c *QueueController) Delqueue() {

}

// @router /queue/status [get]
func (c *QueueController) Status() {
	name := c.GetString("name")
	valid := v.Validation{}
	valid.Required(name, "name")
	if valid.HasErrors() {
		ErrorJsonEnd(&c.G_Controller, "has no queue name")
		return
	}

	qm := new(m.QueueModel)
	if ok, message := qm.GetInfo(name); ok {
		if ok, message := qm.CheckStatus(); ok {
			OKJsonEndAddition(&c.G_Controller, message)
		} else {
			ErrorJsonEnd(&c.G_Controller, message)
		}
	} else {
		ErrorJsonEnd(&c.G_Controller, message)
	}
}

// @router /queue/create [post]
func (c *QueueController) CreateQueue() {
	name := c.GetString("name")
	ttl, err := c.GetInt("ttl")
	if nil != err {
		ErrorJsonEnd(&c.G_Controller, "Params invalid , ttl must be an int")
		return
	}
	qm := new(m.QueueModel)
	valid := v.Validation{}
	valid.Required(name, "name")
	valid.Required(ttl, "ttl")
	if valid.HasErrors() {
		ErrorJsonEnd(&c.G_Controller, "Params has errors ...")
		return
	}
	qm.Init(name, ttl)
	if ok, err := qm.Save(); ok {
		OKJsonEnd(&c.G_Controller)
	} else {
		ErrorJsonEnd(&c.G_Controller, err.Error())
	}
}
