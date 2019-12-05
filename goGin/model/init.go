package model

import (
       "fmt"
       "github.com/jinzhu/gorm"
       _"github.com/jinzhu/gorm/dialects/mysql"
       _"github.com/go-sql-driver/mysql"
)

var Db, err1 = gorm.Open("mysql", "username:password@/dbname?charset=utf8&parseTime?&loc=Local")

func Init() {
     if err1 != nil {
	fmt.Println(err)
	panic("连接数据库失败")
     }
     defer Db.Close()
}
     
     
