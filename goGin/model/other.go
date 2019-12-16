package model

import "log"

func NewCourseToStudent(newcourse string, studentname string) {
     if !CheckStudent(studentname) {
	return
     }
     var value Student
     value.StudentName = studentname
     value.Course = newcourse
     UpdateStudentInfo(value)
}

func NewClassroomToCourse(newclassroom string, course string) {
     if !CheckCourse(course) {
	return
     }
     var value Course
     value.Course = course
     value.Classroom = newclassroom
     UpdateCourseInfo(value)
     return
}

func StudentCourse(student string) []string {
     var courses []Student
     var allcourse []string
     if err := Db.Self.Model(&Student{}).Where("studentname = ?", student).Find(&courses).Error; err != nil {
	log.Println(err)
	panic(err)
     }
     for _, value := range courses {
	 allcourse = append(allcourse, value.Course)
     }
     return allcourse
}

func StudentClassroom(student string) []string {
     var classroom Course
     var courses []Student
     var allclassroom []string
     if err := Db.Self.Model(&Student{}).Where("studentname = ?", student).Find(&courses).Error; err != nil {
	log.Println(err)
	panic(err)
     }
     for _, value := range courses {
	 if err := Db.Self.Model(&Course{}).Where("course = ?", value.Course).First(&classroom).Error; err != nil {
            log.Println(err)
	    panic(err)
	 }
	 allclassroom = append(allclassroom, classroom.Classroom)
     }
     return allclassroom
}



	
	

        
	
