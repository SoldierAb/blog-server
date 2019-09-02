// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)



var(
	adminController =  &controllers.AdminController{}
    markdownController = &controllers.MarkDownController{}
)



func init() {
	beego.InsertFilter("*",beego.BeforeRouter,cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type", "access_token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	ns := beego.NewNamespace("/blog")
	{
		ns.Post("/login",adminController.Login)

		markdownNS := beego.NewNamespace("/markdown")

		markdownNS.Filter("before",adminController.Authentication)
		{
			markdownNS.Post("/add",markdownController.AddMarkDown)
		}

		ns.Namespace(markdownNS)

	}
	beego.AddNamespace(ns)
}
