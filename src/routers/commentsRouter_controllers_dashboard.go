package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["controllers/dashboard:UwCounterController"] = append(beego.GlobalControllerRouter["controllers/dashboard:UwCounterController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/dashboard:UwCounterController"] = append(beego.GlobalControllerRouter["controllers/dashboard:UwCounterController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/dashboard:UwCounterController"] = append(beego.GlobalControllerRouter["controllers/dashboard:UwCounterController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/dashboard:UwCounterController"] = append(beego.GlobalControllerRouter["controllers/dashboard:UwCounterController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/dashboard:UwCounterController"] = append(beego.GlobalControllerRouter["controllers/dashboard:UwCounterController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/dashboard:UwNavController"] = append(beego.GlobalControllerRouter["controllers/dashboard:UwNavController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/dashboard:UwNavController"] = append(beego.GlobalControllerRouter["controllers/dashboard:UwNavController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/dashboard:UwNavController"] = append(beego.GlobalControllerRouter["controllers/dashboard:UwNavController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/dashboard:UwNavController"] = append(beego.GlobalControllerRouter["controllers/dashboard:UwNavController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/dashboard:UwNavController"] = append(beego.GlobalControllerRouter["controllers/dashboard:UwNavController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/dashboard:UwNavListController"] = append(beego.GlobalControllerRouter["controllers/dashboard:UwNavListController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/dashboard:UwNavListController"] = append(beego.GlobalControllerRouter["controllers/dashboard:UwNavListController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/dashboard:UwNavListController"] = append(beego.GlobalControllerRouter["controllers/dashboard:UwNavListController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/dashboard:UwNavListController"] = append(beego.GlobalControllerRouter["controllers/dashboard:UwNavListController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/dashboard:UwNavListController"] = append(beego.GlobalControllerRouter["controllers/dashboard:UwNavListController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/dashboard:UwMonitorController"] = append(beego.GlobalControllerRouter["controllers/dashboard:UwMonitorController"],
		beego.ControllerComments{
			Method: "Alarm",
			Router: `/alarm`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/dashboard:UwMonitorController"] = append(beego.GlobalControllerRouter["controllers/dashboard:UwMonitorController"],
		beego.ControllerComments{
			Method: "HgGraph",
			Router: `/hgGraph`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
