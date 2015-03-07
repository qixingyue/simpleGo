package controllers

import (
	"fmt"
	v "github.com/astaxie/beego/validation"
	m "leafApi/models"
)

type TaskController struct {
	G_Controller
}

func ErrorsToString(errors []*v.ValidationError) string {
	message := ""
	for _, err := range errors {
		message = fmt.Sprintf("%s%s", message, fmt.Sprintf("Name: %s , ErrorMessage : %s \n", err.Key, err.Message))
	}
	return message
}

// @router /task/add [post]
func (c *TaskController) AddTask() {
	uniqueId, queueName, jsonData := c.GetString("uniqueId"), c.GetString("queueName"), c.GetString("jsonData")
	tm := new(m.TaskModel)
	valid := v.Validation{}
	valid.Required(uniqueId, "uniqueId")
	valid.Required(queueName, "queueName")
	valid.Required(jsonData, "jsonData")
	if valid.HasErrors() {
		ErrorJsonEnd(&c.G_Controller, fmt.Sprintf("Param has errors ... %#v ", ErrorsToString(valid.Errors)))
		return
	}
	if ok, message := tm.Init(uniqueId, queueName, jsonData); ok {
		if ok, message := tm.WriteRedis(); ok {
			if tm.ReturnJson != "" {
				OKJsonEndAddition(&c.G_Controller, tm.ReturnJson)
			} else {
				OKJsonEnd(&c.G_Controller)
			}
			return
		} else {
			ErrorJsonEnd(&c.G_Controller, message)
			return
		}
	} else {
		ErrorJsonEnd(&c.G_Controller, message)
		return
	}
}
