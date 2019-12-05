package model

import "github.com/MitsuhaOma/goproject/goGin/model"

type Classroom struct {
     Courses    string
     Roomnumber string
}

func CheckClassroom(classroom Classroom) bool {
     err := Db.Where("roomnumber = ?", classroom.Roomnumber).Error
     return err == nil
}

func CreateClassroom(classroom Classroom) bool {
     Db.create(&classroom)
}


