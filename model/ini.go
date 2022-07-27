package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func Database(connstring string) {
	db, err := gorm.Open("mysql", connstring)
	if err != nil {
		fmt.Println(err)
		panic("Mysql数据库连接错误。")
	}
	fmt.Println("数据库链接成功。")
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)

	}
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(20)
	db.DB().SetMaxIdleConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db
	migrateion()
}
