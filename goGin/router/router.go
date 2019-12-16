package router

import (
	"github.com/gin-gonic/gin"
	"github.com/MitsuhaOma/goproject/goGin/handler"
)

var Router *gin.Engine

func InitRouter() {
     Router = gin.Default()
     Router.POST("/student/create", handler.CreateStudent)
     Router.POST("/classroom/create", handler.CreateClassroom)
     Router.POST("/course/create", handler.CreateCourse)
     Router.POST("/student/assignment", handler.NewCourseToStudent)
     Router.POST("/student/allcourses", handler.StudentCourse)
     Router.POST("/student/allclassrooms", handler.StudentClassroom)


     return
}
