package routers

import (
	"learnbeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//支持Restful

	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.TestController{})

	// 指定匹配方法

	//所有的类型请求均匹配
	beego.Router("/test/add",&controllers.TestController{},"*:AddTest")
	//只匹配get请求
	beego.Router("/test/list",&controllers.TestController{},"get:ListTest")
	//只匹配delete
	beego.Router("/test/delete",&controllers.TestController{},"delete:DeleteTest")

	beego.Include(&controllers.AnnotationController{})




}
