package model

import (
       "github.com/MitsuhaOma/goproject/goGin/model"
)

type Student struct {
     Username string
     Password string
}

func CheckUserByUsername(student Student) bool {
     err := Db.Where("name = ?", student.Username).Error
     return err == nil
}

func CheckPasswordValidate(student Student) bool {
     err := Db.Where("name = ? AND password = ?", student.Username, student.Password).Error
     return err == nil
}

func CreateUser(student Student) {
     Db.Create(&student)
}

