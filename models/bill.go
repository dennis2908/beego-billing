package models

type C_bill struct {
	Id     int    `form:"-"`
	Company_code string `form:"company_code,text,company_code:" valid:"MinSize(5);MaxSize(20)"`
	Company_name string `form:"company_name,text,company_name:"`
	Note string `form:"note,text,address:"`
	Bill int `form:"-"`
}

func (a *C_bill) TableName() string {
	return "c_bill"
}