package controllers

import (
	_ "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/validation"
	models "Learning-Beego/models"
	"strconv"
)

type ManageController struct {
	beego.Controller
}


func (manage *ManageController) View() {
	manage.Layout = "basic-layout.tpl"
	manage.LayoutSections = make(map[string]string)
	manage.LayoutSections["Header"] = "header.tpl"
	manage.LayoutSections["Footer"] = "footer.tpl"
	manage.TplName = "company/view.tpl"
}

func (manage *ManageController) Company_View() {
	o := orm.NewOrm()
	o.Using("default")
	var companies []*models.Company
	var td string
	
	num, err := o.QueryTable("companies").All(&companies)

	if err != orm.ErrNoRows && num > 0 {
		manage.Data["records"] = companies
		for i := 0; i<len(companies);i++ {
			td += "<tr><th>"+strconv.Itoa(i+1)+"</th><th>"+"<button onclick='edit_company("+strconv.Itoa(companies[i].Id)+")' class='btn btn-info'>Edit</button> <button class='btn btn-danger' onclick='del_company("+strconv.Itoa(companies[i].Id)+")'>Delete</button></th>"
			td += "<th>"+companies[i].Company_code+"</th><th>"+companies[i].Company_name+"</th><th>"+companies[i].Address+"</th>" 
			td += "<th>"+companies[i].Phone_number+"</th></tr>" 
		}
		manage.Data["json"] = td
	
	}

	manage.ServeJSON()

}

func (manage *ManageController) Company_API() {
	o := orm.NewOrm()
	o.Using("default")
	var companies []*models.Company
	var td,sql string
	sql = "select * from companies WHERE company_code LIKE '%"+manage.Ctx.Input.Param(":id")+"%' or company_name LIKE '%"+manage.Ctx.Input.Param(":id")+"%'"
	num, err := o.Raw(sql).QueryRows(&companies)
	if err != orm.ErrNoRows && num > 0 {
		manage.Data["records"] = companies
		for i := 0; i<len(companies);i++ {
			td += "<tr><th>"+strconv.Itoa(i+1)+"</th><th>"+"<button onclick='edit_company("+strconv.Itoa(companies[i].Id)+")' class='btn btn-info'>Edit</button> <button class='btn btn-danger' onclick='del_company("+strconv.Itoa(companies[i].Id)+")'>Delete</button></th>"
			td += "<th>"+companies[i].Company_code+"</th><th>"+companies[i].Company_name+"</th><th>"+companies[i].Address+"</th>" 
			td += "<th>"+companies[i].Phone_number+"</th></tr>" 
		}
		
	
	}
	manage.Data["json"] = td

	manage.ServeJSON()

}

func (manage *ManageController) Company_Edit() {
	o := orm.NewOrm()
	o.Using("default")
	var companies []*models.Company
	var sql string
	sql = "select * from companies WHERE id = '"+manage.Ctx.Input.Param(":id")+"'"
	num, err := o.Raw(sql).QueryRows(&companies)
	if err != orm.ErrNoRows && num > 0 {
		manage.Data["json"] = companies[0]		
	}
	
	manage.ServeJSON()

}

func (manage *ManageController) Company_Del() {
	o := orm.NewOrm()
	o.Using("default")
	var companies []*models.Company
	var sql string
	sql = "DELETE FROM companies WHERE id = '"+manage.Ctx.Input.Param(":id")+"'"
	num, err := o.Raw(sql).QueryRows(&companies)
	if err != orm.ErrNoRows && num > 0 {
		manage.Data["json"] = sql		
	}
	
	manage.ServeJSON()

}

func (manage *ManageController) Company_Save() {
    o := orm.NewOrm()
	o.Using("default")
    var company_name,address,phone_number,company_code,id,sql string
	var companies []*models.Company
	manage.Ctx.Input.Bind(&company_code, "Company_code")
	manage.Ctx.Input.Bind(&company_name, "Company_name")
	manage.Ctx.Input.Bind(&address, "Address")
	manage.Ctx.Input.Bind(&phone_number, "Phone_number")
	manage.Ctx.Input.Bind(&id, "Id")
	
	if len(id) > 0 {
		sql = "UPDATE companies SET company_code = '"+company_code+"',company_name = '"+company_name+"',address = '"+address+"',phone_number = '"+phone_number+"' where id = '"+id+"'"
	} else {
		sql = "INSERT INTO companies (company_code, company_name, address, phone_number) VALUES ('"+company_code+"', '"+company_name+"', '"+address+"', '"+phone_number+"')" 
	}
	
	num, err := o.Raw(sql).QueryRows(&companies)
	
	if err != orm.ErrNoRows && num > 0 {
		manage.Data["json"] = sql		
	}
	manage.Data["json"] = len(id)
    manage.ServeJSON()
}
