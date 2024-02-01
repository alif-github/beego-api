// @APIVersion 1.0.0
// @Title Test Learn Basic Beego
// @Description Beego Has a Very Cool Tools to Autogenerate Documents For Your API
// @Contact aysyahputra0215@pnm.co.id
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"test_api/controllers"
	"test_api/middleware"
)

func init() {
	middlewares() //--- Middleware
	ns := beego.NewNamespace("/v1",
		beego.NSBefore(middleware.AuthMiddleware),
		beego.NSNamespace("/user", beego.NSInclude(&controllers.UserController{})),
		beego.NSNamespace("/login", beego.NSInclude(&controllers.LoginController{})),
	)
	beego.AddNamespace(ns)
}

func middlewares() {
	beego.InsertFilter("*", beego.BeforeRouter, middleware.Middleware)
}
