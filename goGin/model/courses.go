package model

import "github.com/MitsuhaOma/goproject/goGin/model"

type Course struct {
     Class string
     Time  string
     Room  string
}

func CheckCourse(course Course) bool {
     err := Db.Where("class = ? AND time = ? AND room = ?", course.Class, course.Time, course.Room).Error
     return err == nil
}

func CreateCourse(course Course) {
     Db.Create(&course)
}
