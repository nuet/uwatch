package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["controllers/configure:UwGraphController"] = append(beego.GlobalControllerRouter["controllers/configure:UwGraphController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/configure:UwGraphController"] = append(beego.GlobalControllerRouter["controllers/configure:UwGraphController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/configure:UwGraphController"] = append(beego.GlobalControllerRouter["controllers/configure:UwGraphController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/configure:UwGraphController"] = append(beego.GlobalControllerRouter["controllers/configure:UwGraphController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/configure:UwGraphController"] = append(beego.GlobalControllerRouter["controllers/configure:UwGraphController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})
}