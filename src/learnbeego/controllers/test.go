package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

func (c *TestController)Get() {
	c.Data["Website"] = "www.taobao.com"
	c.Data["Email"] = "get@163.com"
	c.TplName = "index.tpl"
}

func (c *TestController)Post() {
	c.Data["Website"] = "www.taobao.com"
	c.Data["Email"] = "post@163.com"
	c.TplName = "index.tpl"
}

func (c *TestController)AddTest() {
	c.Data["Website"] = "www.taobao.com"
	c.Data["Email"] = "AddTest@163.com"
	c.TplName = "index.tpl"
}

func (c *TestController)ListTest() {
	c.Data["Website"] = "www.taobao.com"
	c.Data["Email"] = "ListTest@163.com"
	c.TplName = "index.tpl"
}
func (c *TestController)DeleteTest() {
	c.Data["Website"] = "www.taobao.com"
	c.Data["Email"] = "DeleteTest@163.com"
	c.TplName = "index.tpl"
}