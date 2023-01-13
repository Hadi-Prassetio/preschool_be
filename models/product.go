package models

import "time"

type Product struct {
	ID        int       `json:"id"  gorm:"primary_key:auto_increment"`
	Title     string    `json:"title" gorm:"type: varchar(255)"`
	Brand     string    `json:"brand" gorm:"type: varchar(255)"`
	Image     string    `json:"image" gorm:"type: varchar(255)"`
	Price     int       `json:"price" gorm:"type:int"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type ProductUser struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Brand string `json:"brand"`
	Image string `json:"image"`
	Price int    `json:"price"`
}

func (ProductUser) TableName() string {
	return "products"
}
