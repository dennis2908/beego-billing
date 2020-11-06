package controllers

import (
	_ "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/shopspring/decimal"
    "github.com/leekchan/accounting"
	_ "github.com/astaxie/beego/validation"
	models "Learning-Beego/models"
	"strconv"
)

type BillsController struct {
	beego.Controller
}


func (bills *BillsController) View() {
	bills.Layout = "basic-layout.tpl"
	bills.LayoutSections = make(map[string]string)
	bills.LayoutSections["Header"] = "header.tpl"
	bills.LayoutSections["Footer"] = "footer.tpl"
	bills.TplName = "bill/view.tpl"

	o := orm.NewOrm()
	o.Using("default")
	var sql string
	var companies []*models.Company
	sql = "select * from companies order by company_code ASC"
	num, err := o.Raw(sql).QueryRows(&companies)

	if err != orm.ErrNoRows && num > 0 {
		bills.Data["records"] = companies
	}
}

func (bills *BillsController) Bill_View() {
    ac := accounting.Accounting{}
 	o := orm.NewOrm()
	o.Using("default")
	var c_bill []*models.C_bill
	var td,sql string
	var total int
	sql = "select company_bills.id,company_bills.company_code,company_name,bill,note from company_bills join companies on companies.company_code = company_bills.company_code order by company_bills.id DESC limit 25"
	num, err := o.Raw(sql).QueryRows(&c_bill)

	if err != orm.ErrNoRows && num > 0 {
		bills.Data["records"] = c_bill
		for i := 0; i<len(c_bill);i++ {
		    total += c_bill[i].Bill
			td += "<tr><th>"+strconv.Itoa(i+1)+"</th><th>"+"<button onclick='edit_company("+strconv.Itoa(c_bill[i].Id)+")' class='btn btn-info'>Edit</button> <button class='btn btn-danger' onclick='del_company("+strconv.Itoa(c_bill[i].Id)+")'>Delete</button></th>"
			td += "<th>"+c_bill[i].Company_code+" - "+c_bill[i].Company_name+"</th><th>"+c_bill[i].Note+"</th><th style='text-align:right'>"+ac.FormatMoney(c_bill[i].Bill)+"</th></tr>" 
		}
		td += "<tr><th style='text-align:right' colspan=5>Total = "+ac.FormatMoney(total)+"</th></tr>"
		bills.Data["json"] = td
	
	}

	bills.ServeJSON()

}

func (bills *BillsController) Bill_API() {
	ac := accounting.Accounting{}
 	o := orm.NewOrm()
	o.Using("default")
	var c_bill []*models.C_bill
	var td,sql string
	var total int
	sql = "select company_bills.id,company_bills.company_code,company_name,bill,note from company_bills join companies on companies.company_code = company_bills.company_code where company_bills.company_code LIKE '%"+bills.Ctx.Input.Param(":id")+"%' or company_name LIKE '%"+bills.Ctx.Input.Param(":id")+"%' order by company_bills.id DESC limit 25"
	num, err := o.Raw(sql).QueryRows(&c_bill)

	if err != orm.ErrNoRows && num > 0 {
		bills.Data["records"] = c_bill
		for i := 0; i<len(c_bill);i++ {
		    total += c_bill[i].Bill
			td += "<tr><th>"+strconv.Itoa(i+1)+"</th><th>"+"<button onclick='edit_company("+strconv.Itoa(c_bill[i].Id)+")' class='btn btn-info'>Edit</button> <button class='btn btn-danger' onclick='del_company("+strconv.Itoa(c_bill[i].Id)+")'>Delete</button></th>"
			td += "<th>"+c_bill[i].Company_code+" - "+c_bill[i].Company_name+"</th><th>"+c_bill[i].Note+"</th><th style='text-align:right'>"+ac.FormatMoney(c_bill[i].Bill)+"</th></tr>" 
		}
		td += "<tr><th style='text-align:right' colspan=5>Total = "+ac.FormatMoney(total)+"</th></tr>"
		bills.Data["json"] = td
	
	}
	//bills.Data["json"] = sql

	bills.ServeJSON()

}

func (bills *BillsController) Bill_Edit() {
	o := orm.NewOrm()
	o.Using("default")
	var c_bill []*models.C_bill
	var sql string
	sql = "select * from company_bills WHERE id = '"+bills.Ctx.Input.Param(":id")+"'"
	num, err := o.Raw(sql).QueryRows(&c_bill)
	if err != orm.ErrNoRows && num > 0 {
		bills.Data["json"] = c_bill[0]		
	}
	
	bills.ServeJSON()

}

func (bills *BillsController) Bill_Del() {
	o := orm.NewOrm()
	o.Using("default")
	var c_bill []*models.C_bill
	var sql string
	sql = "DELETE FROM company_bills WHERE id = '"+bills.Ctx.Input.Param(":id")+"'"
	num, err := o.Raw(sql).QueryRows(&c_bill)
	if err != orm.ErrNoRows && num > 0 {
		bills.Data["json"] = sql		
	}
	
	bills.ServeJSON()

}

func (bills *BillsController) Bill_Save() {
    o := orm.NewOrm()
	o.Using("default")
    var company_code,bill,note,id,sql string
	var c_bill []*models.C_bill
	bills.Ctx.Input.Bind(&company_code, "Company_code")
	bills.Ctx.Input.Bind(&bill, "Bill")
	bills.Ctx.Input.Bind(&note, "Note")
	bills.Ctx.Input.Bind(&id, "Id")
	
	if len(id) > 0 {
		sql = "UPDATE company_bills SET company_code = '"+company_code+"',bill = '"+bill+"',note = '"+note+"' where id = '"+id+"'"
	} else {
		sql = "INSERT INTO company_bills (company_code, bill, note) VALUES ('"+company_code+"', '"+bill+"', '"+note+"')" 
	}
	
	num, err := o.Raw(sql).QueryRows(&c_bill)
	
	if err != orm.ErrNoRows && num > 0 {
		bills.Data["json"] = sql		
	}
	bills.Data["json"] = len(id)
    bills.ServeJSON()
}