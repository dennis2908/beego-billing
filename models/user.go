package models

type User struct {
	Id           int    `orm:"auto"`
	Username     string `orm:"unique;size(100)"`
	PasswordHash string `orm:"size(255)"`
}

func (user *User) TableName() string {
	return "users"
}
