package models

import "time"

type Admin struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	FullName  string    `json:"fullname" gorm:"type: varchar(100)"`
	AdminName string    `json:"adminname" gorm:"type: varchar(100)"`
	Email     string    `json:"email" gorm:"type: varchar(70)"`
	Phone     string    `json:"phone" gorm:"type: varchar(15)"`
	Password  string    `json:"password" gorm:"type: varchar(100)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type AdminProfile struct {
	ID       int    `json:"id"  `
	FullName string `json:"fullname" `
	Email    string `json:"email" `
}

func (AdminProfile) TableName() string {
	return "admins"
}