package models

import "time"

type Visitor struct {
	ID        int          `json:"id"  gorm:"primary_key:auto_increment"`
	Name      string       `json:"name" gorm:"type: varchar(100)"`
	Email     string       `json:"email" gorm:"type: varchar(100)"`
	Phone     string       `json:"phone" gorm:"type: varchar(15)"`
	ChildName string       `json:"child_name" gorm:"type: varchar(100)"`
	ChildAge  int          `json:"child_age" gorm:"type: integer(2)"`
	Message   string       `json:"message" gorm:"type:varchar(225)"`
	Status    string       `json:"status" gorm:"type:varchar(100)"`
	AdminID   int          `json:"admin_id"`
	Admin     AdminProfile `json:"admin"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
}

type VisitorDetail struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	ChildName string `json:"child_name"`
	ChildAge  int    `json:"child_age"`
	Message   string `json:"message"`
	Status    string `json:"status"`
	AdminID   int    `json:"admin_id"`
}

func (VisitorDetail) TableName() string {
	return "visitors"
}
