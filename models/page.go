package models

import "time"

type Page struct {
	ID        int          `json:"id"  gorm:"primary_key:auto_increment"`
	Type      string       `json:"type" gorm:"type: varchar(100)"`
	Title     string       `json:"title" gorm:"type: varchar(100)"`
	Desc      string       `json:"desc" gorm:"type: varchar(225)"`
	Email     string       `json:"email" gorm:"type: varchar(100)"`
	Phone     string       `json:"phone" gorm:"type: varchar(15)"`
	AdminID   int          `json:"admin_id" gorm:"foreignKey:ClassID" `
	Admin     AdminProfile `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
}

type PageDetail struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func (PageDetail) TableName() string {
	return "pages"
}
type CreatePage struct {
	Type  string `json:"type" validate:"required"`
	Title string `json:"title" validate:"required"`
	Desc  string `json:"desc" validate:"required"`
	Email string `json:"email" validate:"required"`
	Phone string `json:"phone" validate:"required"`
}
type UpdatePage struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
