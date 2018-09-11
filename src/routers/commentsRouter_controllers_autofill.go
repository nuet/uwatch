package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["controllers/autofill:UwAutofillController"] = append(beego.GlobalControllerRouter["controllers/autofill:UwAutofillController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/autofill:UwAutofillController"] = append(beego.GlobalControllerRouter["controllers/autofill:UwAutofillController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/autofill:UwAutofillController"] = append(beego.GlobalControllerRouter["controllers/autofill:UwAutofillController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/autofill:UwAutofillController"] = append(beego.GlobalControllerRouter["controllers/autofill:UwAutofillController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/autofill:UwAutofillController"] = append(beego.GlobalControllerRouter["controllers/autofill:UwAutofillController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/autofill:UwAutofillController"] = append(beego.GlobalControllerRouter["controllers/autofill:UwAutofillController"],
		beego.ControllerComments{
			Method: "Detection",
			Router: `/detection/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/autofill:UwOperationController"] = append(beego.GlobalControllerRouter["controllers/autofill:UwOperationController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/autofill:UwOperationController"] = append(beego.GlobalControllerRouter["controllers/autofill:UwOperationController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/autofill:UwOperationController"] = append(beego.GlobalControllerRouter["controllers/autofill:UwOperationController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/autofill:UwOperationController"] = append(beego.GlobalControllerRouter["controllers/autofill:UwOperationController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/autofill:UwOperationController"] = append(beego.GlobalControllerRouter["controllers/autofill:UwOperationController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/autofill:UwOplogController"] = append(beego.GlobalControllerRouter["controllers/autofill:UwOplogController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/autofill:UwOplogController"] = append(beego.GlobalControllerRouter["controllers/autofill:UwOplogController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/autofill:UwOplogController"] = append(beego.GlobalControllerRouter["controllers/autofill:UwOplogController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/autofill:UwOplogController"] = append(beego.GlobalControllerRouter["controllers/autofill:UwOplogController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/autofill:UwOplogController"] = append(beego.GlobalControllerRouter["controllers/autofill:UwOplogController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/autofill:UwTaskController"] = append(beego.GlobalControllerRouter["controllers/autofill:UwTaskController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/autofill:UwTaskController"] = append(beego.GlobalControllerRouter["controllers/autofill:UwTaskController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers/autofill:UwTaskController"] = append(beego.GlobalControllerRouter["controllers/autofill:UwTaskController"],
		beego.ControllerComments{
			Method: "GetRecords",
			Router: `/getRecords`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
