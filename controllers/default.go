package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (main *MainController) Get() {
	main.Layout = "basic-layout.tpl"
	main.LayoutSections = make(map[string]string)
	main.LayoutSections["Header"] = "header.tpl"
	main.LayoutSections["Footer"] = "footer.tpl"
	main.TplName = "manage/home.tpl"
}

func (main *MainController) HelloSitepoint() {
	main.Data["Website"] = "My Website"
	main.Data["Email"] = "your.email.address@example.com"
	main.Data["EmailName"] = "Your Name"
	main.TplName = "default/hello-sitepoint.tpl"
}
