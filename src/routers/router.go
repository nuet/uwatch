// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"controllers"
	"controllers/autofill"
	"controllers/falcon"
	"controllers/dashboard"
	"controllers/configure"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"time"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		MaxAge:           5 * time.Minute,
		AllowCredentials: true,
	}))

	beego.Router("/", &controllers.BaseRouter{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})

	//open-falcon回调 URL
	beego.Router("/callback", &controllers.MetricsCallBackController{})
	beego.Router("/im_callback", &falcon.AlertCallBackController{},"post:ImCallback")
	beego.Router("/mail_callback", &falcon.AlertCallBackController{},"post:MailCallback")
	beego.Router("/voice_callback", &falcon.AlertCallBackController{},"post:VoiceCallback")
	beego.Router("/test", &falcon.AlertCallBackController{},"get:Test")

	//前端调用接口
	beego.Router("/uw_graph_data", &dashboard.GetGraphDataController{},"get:FindOneGraphData")
	beego.Router("/findFalconData", &dashboard.GetGraphDataController{},"get:FindOneFalconData")
	beego.Router("/findOneFalconData", &dashboard.GetGraphDataController{},"get:GetOneFalconData")
	beego.Router("/getCounter", &dashboard.GetGraphDataController{},"get:GetCounter")
	beego.Router("/getGraphCounter", &dashboard.GetGraphDataController{},"get:GetGraphCounter")
	beego.Router("/getCmdbHost", &dashboard.GetCmdbHostController{},"get:GetCmdbHostList")
	//兼容0.2版本falcon
	beego.Router("/getFalconCounter", &dashboard.GetFalconCounterController{},"get:GetFalconCounter")

	//拉取导航
	beego.Router("/uw_navall", &dashboard.UwNavController{},"get:GetNavAll")

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/uw_autofill",
			beego.NSInclude(
				&autofill.UwAutofillController{},
			),
		),

		beego.NSNamespace("/uw_operation",
			beego.NSInclude(
				&autofill.UwOperationController{},
			),
		),

		beego.NSNamespace("/uw_oplog",
			beego.NSInclude(
				&autofill.UwOplogController{},
			),
		),

		beego.NSNamespace("/uw_task",
			beego.NSInclude(
				&autofill.UwTaskController{},
			),
		),

		beego.NSNamespace("/search_keyword",
			beego.NSInclude(
				&controllers.SearchKeywordController{},
			),
		),

		beego.NSNamespace("/system_role",
			beego.NSInclude(
				&controllers.SystemRoleController{},
			),
		),

		beego.NSNamespace("/system_user",
			beego.NSInclude(
				&controllers.SystemUserController{},
			),
		),

		beego.NSNamespace("/uw_graph",
			beego.NSInclude(
				&configure.UwGraphController{},
			),
		),

		beego.NSNamespace("/uw_counter",
			beego.NSInclude(
				&dashboard.UwCounterController{},
			),
		),

		beego.NSNamespace("/uw_monitor",
			beego.NSInclude(
				&dashboard.UwMonitorController{},
			),
		),

		beego.NSNamespace("/uw_nav",
			beego.NSInclude(
				&dashboard.UwNavController{},
			),
		),

		beego.NSNamespace("/uw_navmenu",
			beego.NSInclude(
				&dashboard.UwNavListController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
