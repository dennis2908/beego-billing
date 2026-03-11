package routers

import (
	"github.com/astaxie/beego"
	"Learning-Beego/controllers"
)

func init() {
	beego.Router("/", &controllers.ManageController{}, "get:View")
    beego.Router("/company", &controllers.ManageController{}, "get:View")
    beego.Router("/company/view/", &controllers.ManageController{}, "get:Company_View")
	beego.Router("/company/view/:id", &controllers.ManageController{}, "get:Company_API")
	beego.Router("/company/edit/:id", &controllers.ManageController{}, "get:Company_Edit")
	beego.Router("/company/delete/:id", &controllers.ManageController{}, "get:Company_Del")
	beego.Router("/company/save/", &controllers.ManageController{}, "get:Company_Save")
	beego.Router("/bill", &controllers.BillsController{}, "get:View")
    beego.Router("/bill/view/", &controllers.BillsController{}, "get:Bill_View")
	beego.Router("/bill/view/:id", &controllers.BillsController{}, "get:Bill_API")
	beego.Router("/bill/edit/:id", &controllers.BillsController{}, "get:Bill_Edit")
	beego.Router("/bill/delete/:id", &controllers.BillsController{}, "get:Bill_Del")
	beego.Router("/bill/save/", &controllers.BillsController{}, "get:Bill_Save")
}
