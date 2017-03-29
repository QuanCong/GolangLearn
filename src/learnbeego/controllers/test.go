package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

func (c *TestController)Get()  {
	c.Data["Website"]= "www.taobao.com"
	c.Data["Email"]="songxu9185@163.com"
	c.TplName = "index.tpl"

}
