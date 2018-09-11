package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["controllers:AssetController"] = append(beego.GlobalControllerRouter["controllers:AssetController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:AssetController"] = append(beego.GlobalControllerRouter["controllers:AssetController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:AssetController"] = append(beego.GlobalControllerRouter["controllers:AssetController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:AssetController"] = append(beego.GlobalControllerRouter["controllers:AssetController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:AssetController"] = append(beego.GlobalControllerRouter["controllers:AssetController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:AssetController"] = append(beego.GlobalControllerRouter["controllers:AssetController"],
		beego.ControllerComments{
			Method: "MultiDelete",
			Router: `/del/multi`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:AssetController"] = append(beego.GlobalControllerRouter["controllers:AssetController"],
		beego.ControllerComments{
			Method: "CheckAsCode",
			Router: `/checkAsCode/:code`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:AssetController"] = append(beego.GlobalControllerRouter["controllers:AssetController"],
		beego.ControllerComments{
			Method: "CopyAsset",
			Router: `/copy/:num`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:AssetController"] = append(beego.GlobalControllerRouter["controllers:AssetController"],
		beego.ControllerComments{
			Method: "UploadAsset",
			Router: `/upload`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:ConsoleScriptController"] = append(beego.GlobalControllerRouter["controllers:ConsoleScriptController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:ConsoleScriptController"] = append(beego.GlobalControllerRouter["controllers:ConsoleScriptController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:ConsoleScriptController"] = append(beego.GlobalControllerRouter["controllers:ConsoleScriptController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:ConsoleScriptController"] = append(beego.GlobalControllerRouter["controllers:ConsoleScriptController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:ConsoleScriptController"] = append(beego.GlobalControllerRouter["controllers:ConsoleScriptController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:RecoveryTsakController"] = append(beego.GlobalControllerRouter["controllers:RecoveryTsakController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:RecoveryTsakController"] = append(beego.GlobalControllerRouter["controllers:RecoveryTsakController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:SearchKeywordController"] = append(beego.GlobalControllerRouter["controllers:SearchKeywordController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:SearchKeywordController"] = append(beego.GlobalControllerRouter["controllers:SearchKeywordController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:SearchKeywordController"] = append(beego.GlobalControllerRouter["controllers:SearchKeywordController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:SearchKeywordController"] = append(beego.GlobalControllerRouter["controllers:SearchKeywordController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:SearchKeywordController"] = append(beego.GlobalControllerRouter["controllers:SearchKeywordController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:SystemRoleController"] = append(beego.GlobalControllerRouter["controllers:SystemRoleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:SystemRoleController"] = append(beego.GlobalControllerRouter["controllers:SystemRoleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:SystemRoleController"] = append(beego.GlobalControllerRouter["controllers:SystemRoleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:SystemRoleController"] = append(beego.GlobalControllerRouter["controllers:SystemRoleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:SystemRoleController"] = append(beego.GlobalControllerRouter["controllers:SystemRoleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:SystemUserController"] = append(beego.GlobalControllerRouter["controllers:SystemUserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:SystemUserController"] = append(beego.GlobalControllerRouter["controllers:SystemUserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:SystemUserController"] = append(beego.GlobalControllerRouter["controllers:SystemUserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:SystemUserController"] = append(beego.GlobalControllerRouter["controllers:SystemUserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:SystemUserController"] = append(beego.GlobalControllerRouter["controllers:SystemUserController"],
		beego.ControllerComments{
			Method: "GetUserList",
			Router: `/getUserList`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:SystemUserController"] = append(beego.GlobalControllerRouter["controllers:SystemUserController"],
		beego.ControllerComments{
			Method: "GetTeamList",
			Router: `/getTeamList`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:SystemUserController"] = append(beego.GlobalControllerRouter["controllers:SystemUserController"],
		beego.ControllerComments{
			Method: "GetDepartmentList",
			Router: `/getDepartmentList`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:SystemUserController"] = append(beego.GlobalControllerRouter["controllers:SystemUserController"],
		beego.ControllerComments{
			Method: "GetList",
			Router: `/getList`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:SystemUserController"] = append(beego.GlobalControllerRouter["controllers:SystemUserController"],
		beego.ControllerComments{
			Method: "GetUserRole",
			Router: `/getUserRole`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
