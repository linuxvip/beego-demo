// @APIVersion 1.0.0
// @Title Devops API
// @Description Devops System APIs.
// @Contact liuchengming@laiye.com

package routers

import (
	"github.com/astaxie/beego"
	"devops/controllers"
)

// 使用注释路由
func init() {

	beego.Router("/", &controllers.DefaultController{}, "*:GetAll")
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
