package model

type User struct {
	Id 				int `json:"id" form:"id" gorm:"primary_key"`
	Uid 			string `json:"uid" form:"uid"`
	Username 		string `json:"username" form:"username"`
	Phone 			string `json:"phone" form:"phone"`
}

func (User) TableName() string {
	return "user"
}