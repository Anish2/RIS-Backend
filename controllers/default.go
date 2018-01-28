package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type BasicResp struct {
	status string `json:"status"`
}

func (c *MainController) Post() {
	user := c.GetString("username")
	pass := c.GetString("password")
	beego.Info(user)
	res := BasicResp{"error"}
	if user == "Anish" && pass == "Anish" {
		res = BasicResp{"success"}
	} else {
		res = BasicResp{"error"}
	}
	c.Data["json"] = &res
	c.ServeJSON()
}
