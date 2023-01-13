package database

import (
	"fmt"
	"ibox/models"
	"ibox/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Transaction{}, &models.Cart{}, &models.Order{})

	if err != nil {
		fmt.Println(err)
		panic("migration error")
	}
	fmt.Println("migration success")
}
