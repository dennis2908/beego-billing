package models

type Company struct {
	Id     int    `form:"-"`
	Company_code string `form:"company_code,text,company_code:" valid:"MinSize(5);MaxSize(20)"`
	Company_name string `form:"company_name,text,company_name:"`
	Address string `form:"address,text,address:"`
	Phone_number string `form:"phone_number,text,url:"`
}

func (a *Company) TableName() string {
	return "companies"
}