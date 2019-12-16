package model

import "log"

type Course struct {
     Course     string `gorm:"name" json:"coursename"`
     Classroom  string `gorm:"classroom" json:"classroom"`
     CourseInfo string `gorm:"course_info" json:"courseinfo"`
}

func CheckCourse(coursename string) bool {
     var tmpcourse Course
     if err := Db.Self.Model(&Course{}).Where("name = ?", coursename).First(&tmpcourse).Error; err != nil {
	log.Println(err)
	return false
     }
     return true
}

func CreateCourse(coursename string) string {
     var tmpcourse Course 
     if err := Db.Self.Model(&Course{}).Where("name = ?", coursename).First(&tmpcourse).Error; err != nil {
	log.Println(err)
	return "数据库链接出现错误"    
     } 
     if CheckCourse(tmpcourse.Course) {
	return "课程" + coursename + "已存在"
     }
     Db.Self.Model(&Course{}).Create(&tmpcourse)
     return "课程" + coursename + "成功创建"
}

func UpdateCourseInfo(course Course) {
     if err := Db.Self.Model(&Course{}).Where("name = ?", course.Course).First(&course).Error; err != nil {
	log.Println(err)
	panic(err)
     }
     return
}

