package main

import (
       "fmt"
       "github.com/jinzhu/gorm"
       _"github.com/jinzhu/gorm/dialects/mysql"
       _"github.com/go-sql-driver/mysql"
)

type User struct {
     Name     string `gorm:"Name"`
     Password string `gorm:"Password"`
}

func main() {
	db, err := gorm.Open("mysql", "root:4399687@/goweb?charset=utf8&parseTime?&loc=Local")
     if err != nil {
	fmt.Println(err)
	panic("连接数据库失败")
     }
     defer db.Close()
     
     db.AutoMigrate(&User{})

     var user User
     db.First(&user)
     fmt.Println(user.Name, user.Password)
     
    
}

