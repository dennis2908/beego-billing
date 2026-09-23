package controllers

import (
	models "Learning-Beego/models"
	"encoding/json"
	_ "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/validation"
	"net/http"
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
		for i := 0; i < len(companies); i++ {
			td += "<tr><th>" + strconv.Itoa(i+1) + "</th><th>" + "<button onclick='edit_company(" + strconv.Itoa(companies[i].Id) + ")' class='btn btn-info'>Edit</button> <button class='btn btn-danger' onclick='del_company(" + strconv.Itoa(companies[i].Id) + ")'>Delete</button></th>"
			td += "<th>" + companies[i].Company_code + "</th><th>" + companies[i].Company_name + "</th><th>" + companies[i].Address + "</th>"
			td += "<th>" + companies[i].Phone_number + "</th></tr>"
		}
		manage.Data["json"] = td

	}

	manage.ServeJSON()

}

func (manage *ManageController) Company_API() {
	o := orm.NewOrm()
	o.Using("default")
	var companies []*models.Company
	var td, sql string
	sql = "select * from companies WHERE company_code LIKE '%" + manage.Ctx.Input.Param(":id") + "%' or company_name LIKE '%" + manage.Ctx.Input.Param(":id") + "%'"
	num, err := o.Raw(sql).QueryRows(&companies)
	if err != orm.ErrNoRows && num > 0 {
		manage.Data["records"] = companies
		for i := 0; i < len(companies); i++ {
			td += "<tr><th>" + strconv.Itoa(i+1) + "</th><th>" + "<button onclick='edit_company(" + strconv.Itoa(companies[i].Id) + ")' class='btn btn-info'>Edit</button> <button class='btn btn-danger' onclick='del_company(" + strconv.Itoa(companies[i].Id) + ")'>Delete</button></th>"
			td += "<th>" + companies[i].Company_code + "</th><th>" + companies[i].Company_name + "</th><th>" + companies[i].Address + "</th>"
			td += "<th>" + companies[i].Phone_number + "</th></tr>"
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
	sql = "select * from companies WHERE id = '" + manage.Ctx.Input.Param(":id") + "'"
	num, err := o.Raw(sql).QueryRows(&companies)
	if err != orm.ErrNoRows && num > 0 {
		manage.Data["json"] = companies[0]
	}

	manage.ServeJSON()

}

func (manage *ManageController) Company_Del() {
	o := orm.NewOrm()
	o.Using("default")
	id := manage.Ctx.Input.Param(":id")
	if id == "" {
		manage.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		manage.Data["json"] = map[string]interface{}{"success": false, "message": "Company ID is required"}
		manage.ServeJSON()
		return
	}

	sql := "DELETE FROM companies WHERE id = '" + id + "'"
	result, err := o.Raw(sql).Exec()
	if err != nil {
		manage.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		manage.Data["json"] = map[string]interface{}{"success": false, "message": "Failed to delete company", "error": err.Error()}
		manage.ServeJSON()
		return
	}
	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		manage.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		manage.Data["json"] = map[string]interface{}{"success": false, "message": "Company not found"}
		manage.ServeJSON()
		return
	}

	manage.Data["json"] = map[string]interface{}{"success": true, "message": "Successfully deleted company"}
	manage.ServeJSON()
}

func (manage *ManageController) Company_Save() {
	o := orm.NewOrm()
	o.Using("default")
	var company_name, address, phone_number, company_code, id, sql string
	if len(manage.Ctx.Input.RequestBody) > 0 {
		var payload struct {
			CompanyCode string `json:"Company_code"`
			CompanyName string `json:"Company_name"`
			Address     string `json:"Address"`
			PhoneNumber string `json:"Phone_number"`
			ID          string `json:"Id"`
		}
		if err := json.Unmarshal(manage.Ctx.Input.RequestBody, &payload); err != nil {
			manage.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			manage.Data["json"] = map[string]interface{}{"success": false, "message": "Invalid JSON request"}
			manage.ServeJSON()
			return
		}
		company_code = payload.CompanyCode
		company_name = payload.CompanyName
		address = payload.Address
		phone_number = payload.PhoneNumber
		id = payload.ID
	} else {
		manage.Ctx.Input.Bind(&company_code, "Company_code")
		manage.Ctx.Input.Bind(&company_name, "Company_name")
		manage.Ctx.Input.Bind(&address, "Address")
		manage.Ctx.Input.Bind(&phone_number, "Phone_number")
		manage.Ctx.Input.Bind(&id, "Id")
	}

	duplicateSQL := "SELECT id FROM companies WHERE company_code = ?"
	duplicateArgs := []interface{}{company_code}
	if len(id) > 0 {
		duplicateSQL += " AND id <> ?"
		duplicateArgs = append(duplicateArgs, id)
	}
	var duplicateID int
	duplicateErr := o.Raw(duplicateSQL, duplicateArgs...).QueryRow(&duplicateID)
	if duplicateErr == nil {
		manage.Ctx.ResponseWriter.WriteHeader(http.StatusConflict)
		manage.Data["json"] = map[string]interface{}{"success": false, "message": "Company code already exists"}
		manage.ServeJSON()
		return
	}
	if duplicateErr != orm.ErrNoRows {
		manage.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		manage.Data["json"] = map[string]interface{}{"success": false, "message": "Failed to check company code", "error": duplicateErr.Error()}
		manage.ServeJSON()
		return
	}

	if len(id) > 0 {
		sql = "UPDATE companies SET company_code = '" + company_code + "',company_name = '" + company_name + "',address = '" + address + "',phone_number = '" + phone_number + "' where id = '" + id + "'"
	} else {
		sql = "INSERT INTO companies (company_code, company_name, address, phone_number) VALUES ('" + company_code + "', '" + company_name + "', '" + address + "', '" + phone_number + "')"
	}

	result, err := o.Raw(sql).Exec()
	if err != nil {
		manage.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		manage.Data["json"] = map[string]interface{}{"success": false, "message": "Failed to save or update company", "error": err.Error()}
		manage.ServeJSON()
		return
	}

	affected, err := result.RowsAffected()
	if err != nil || (len(id) > 0 && affected == 0) {
		manage.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		manage.Data["json"] = map[string]interface{}{"success": false, "message": "Company not found or unchanged"}
		manage.ServeJSON()
		return
	}

	message := "Successfully saved company"
	if len(id) > 0 {
		message = "Successfully updated company"
	}
	manage.Data["json"] = map[string]interface{}{"success": true, "message": message}
	manage.ServeJSON()
}
