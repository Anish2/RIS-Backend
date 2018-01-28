package controllers

import (
	"encoding/base64"
	"strings"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type BasicResp struct {
	Status string `json:"status"`
}

type BlockData struct {
	Name           string `json:"name"`
	Birthplace     string `json:"birthplace"`
	MedicalHistory string `json:"medicalhistory"`
	Locations      string `json:"locations"`
	GeneralHealth  string `json:"health"`
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

type IdentifyController struct {
	beego.Controller
}

func (c *IdentifyController) Post() {
	data := c.GetString("data")
	sDec, err := base64.StdEncoding.DecodeString(strings.Split(data, ",")[1])
	beego.Debug(data)
	if err != nil {
		beego.Error(err)
	}
	// 1) Detect
	// 2) Identify
	// 3a) Add Person, Train
	// 3b) Get Block
	// 4) Post Block
	faceId, err := DetectFace(&sDec)
	if err != nil {
		beego.Error(err)
	}
	personId, err := IdentifyFace(faceId)
	if err != nil {
		beego.Error(err)
		res := BasicResp{"error"}
		c.Data["json"] = &res
		c.ServeJSON()
	}
	blockId, err := GetPerson(personId)
	if err != nil {
		beego.Error(err)
	}
	blockData, err := GetBlock(blockId)
	if err != nil {
		beego.Error(err)
	}
	beego.Debug(blockData["medicalhistory"])
	c.Data["json"] = BlockData{blockData["name"].(string), blockData["birthplace"].(string), blockData["medicalhistory"].(string), blockData["locations"].(string), blockData["health"].(string)}
	c.ServeJSON()
}

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Post() {
	image := c.GetString("image")
	sDec, err := base64.StdEncoding.DecodeString(strings.Split(image, ",")[1])
	if err != nil {
		beego.Error(err)
	}
	blockId, err := PostBlock(c)
	personId := NewPerson(c.GetString("name"), blockId)
	AddFace(personId, &sDec)
	Train()
	res := BasicResp{"success"}
	c.Data["json"] = &res
	c.ServeJSON()
}
