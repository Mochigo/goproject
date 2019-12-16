package model

import "log"

type Classroom struct {
     Classroom  string `gorm:"classroom" json:"classroom"`
}

func CheckClassroom(classroom string) bool {
     var newclassroom Classroom
     if err := Db.Self.Model(&Classroom{}).Where("classroom = ?",classroom).First(&newclassroom).Error; err != nil {
	log.Println(err)
	return false
     }
     return true
}

func CreateClassroom(classroom Classroom) {
     if CheckClassroom(classroom.Classroom) {
	return
     }
     Db.Self.Model(&Classroom{}).Create(&classroom)
}
     


