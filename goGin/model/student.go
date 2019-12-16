package model

import "log"

type Student struct {
     StudentName `gorm: "studentname" json: "studentname"`
     Course   `gorm:"course" json:"course"`
}

func CheckStudent(studentname string) bool {
     var tmpstudent Student
     if err := Db.Self.Model(&Student{}).Where("studentname = ?", studentname).First(&tmpstudent); len(tmpstudent.StudentName) != 0 {
        return true
     }
     return false
}

func CreateStudent(studentname string) string {
     if CheckStudent(studentname) {
	return "学生"+studentname+"已存在"
     }
     var tmpstudent = Student{StudentName:studentname}
     if err := Db.Self.Model(&Student{}).Create(&tmpstudent).Error; err != nil {
	log.Println(err)
	return "创建失败"
     }
     return "创建成功"
}
 
func UpdateStudentInfo(student Student) {
     if err := Db.Self.Model(&Student{}).Where("studentname = ?", student.StudentName).Update(&student).Error; err != nil {
	log.Println(err)
	return
     }
}
