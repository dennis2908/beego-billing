package routers

import (
	"Learning-Beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, controllers.AuthFilter, true)
	beego.Router("/login", &controllers.AuthController{}, "get:Login;post:Login")
	beego.Router("/logout", &controllers.AuthController{}, "get:Logout;post:Logout")
	beego.Router("/user", &controllers.AuthController{}, "post:User_Save;get:CurrentUser")
	beego.Router("/", &controllers.ManageController{}, "get:View")
	beego.Router("/company", &controllers.ManageController{}, "get:View")
	beego.Router("/company/view/", &controllers.ManageController{}, "get:Company_View")
	beego.Router("/company/view/:id", &controllers.ManageController{}, "get:Company_API")
	beego.Router("/company/edit/:id", &controllers.ManageController{}, "get:Company_Edit")
	beego.Router("/company/delete/:id", &controllers.ManageController{}, "delete:Company_Del")
	beego.Router("/company/save/", &controllers.ManageController{}, "post,put:Company_Save")
	beego.Router("/bill", &controllers.BillsController{}, "get:View")
	beego.Router("/bill/view/", &controllers.BillsController{}, "get:Bill_View")
	beego.Router("/bill/view/:id", &controllers.BillsController{}, "get:Bill_API")
	beego.Router("/bill/edit/:id", &controllers.BillsController{}, "get:Bill_Edit")
	beego.Router("/bill/delete/:id", &controllers.BillsController{}, "delete:Bill_Del")
	beego.Router("/bill/save/", &controllers.BillsController{}, "post,put:Bill_Save")
}
