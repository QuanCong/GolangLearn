package controllers

import "github.com/astaxie/beego"

type AnnotationController struct {
	beego.Controller //嵌入beego.Controller
}
//添加映射
//用户如果没有进行注册，那么就会通过反射来执行对应的函数，
// 如果注册了就会通过 interface 来进行执行函数，性能上面会提升很多。
func (c *AnnotationController)URLMapping() {

	/*
	 * 注意 此处第一个参数应当是对应rest的方法的驼峰写法
	 * 并且相同的前缀只能出现一次，如List注册后，不可以有ListAdd，ListDel之类的
	 */
	c.Mapping("List", c.List)
	c.Mapping("AddList", c.ListAdd)
	c.Mapping("GetList", c.ListAll)
}



//@router /annotation/list [get]
func (c *AnnotationController)List() {
	c.Data["Website"] = "AnnotationController"
	c.Data["Email"] = "List@163.com"
	c.TplName = "index.tpl"
}

//@router /annotation/getList [get]
func (c *AnnotationController)ListAll() {
	c.Data["Website"] = "AnnotationController"
	c.Data["Email"] = "getList@163.com"
	c.TplName = "index.tpl"
}
//@router /annotation/addList [get]
func (c *AnnotationController)ListAdd() {
	c.Data["Website"] = "AnnotationController"
	c.Data["Email"] = "listAdd@163.com"
	c.TplName = "index.tpl"
}

