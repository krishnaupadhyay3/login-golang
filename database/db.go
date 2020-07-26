package db

import (
	"blogapp/model"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetPostGreSql() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", PgsQLHost)
	if err != nil {
		fmt.Println(err)
		panic("Cannot Connect To DataBase")
	}
	db.AutoMigrate(&model.User{})
	return db, nil
}
