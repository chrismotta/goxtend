package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

// Connect connects to db
func Connect() *gorm.DB {
	db, err = gorm.Open("mysql", "root:123@/jazmeendeco?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
