package controllers

import (
	"os"

	"github.com/astaxie/beego"
	"github.com/grayzone/three/util"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
	c.Layout = "layout.html"
}

func (c *MainController) LoadData() {
	filename := c.GetString("filename")
	pwd, _ := os.Getwd()
	fullpath := pwd + "/data/" + filename
	beego.Info("full path :", fullpath)
	//200 x 160 x 160
	result := util.LoadMRIRaw(fullpath, 200, 160, 160)
	c.Data["json"] = &result
	c.ServeJSON()

}
