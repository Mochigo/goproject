package handler

import(
      "github.com/gin-gonic/gin"
      "github.com/MitsuhaOma/goproject/goGin/model"
)

type students struct {
     Student []string `json:"student"`
}

func CreateStudent(c *gin.Context) {
     var tmpStudent students
     if err := c.BindJSON(&tmpStudent); err != nil {
	c.JSON(400,gin.H{
	      "message":"Bad Request!",
        })
	return
     }
     for _,v := range tmpStudent.Student {
	 c.JSON(200,gin.H{
		"message":model.CreateStudent(v),
	 })
     }
     c.JSON(200,gin.H{
	    "message":"Creat Success",
     })
}

func NewCourseToStudent(c *gin.Context) {
     var tmpstudent Student
     if err := c.BindJSON(&tmpstudent);err != nil {
	c.JSON(400, gin.H{
	      "message":"Bad Request!",
	})
	return
     }
     NewCourseToStudent(tmpstudent.Course, tmpstudent.StudentName)
     c.JSON(200, gin.H{
	   "message":"Success!",
     })
}

func StudentCourse(c *gin.Context) {
     var course students
     if err := c.BindJSON(&course); err != nil {
	c.JSON(400, gin.H{
	      "message": "Bad Request!",
	})
	return
     }
     for _, value := range course.Student {
	 c.JSON(200, gin.H{
		"message": model.StudentCourse(value),
         })
     }
     c.JSON(200, gin.H{
           "message":"Query Success",
     })
}

func StudentClassroom(c *gin.Context) {
     var classroom students
     if err := c.BindJSON(&classroom); err != nil {
	c.JSON(400, gin.H{
	      "message": "Bad Request!",
	})
	return
     }
     for _, value := range classroom.Student {
	 c.JSON(200, gin.H{
		"message": model.StudentClassroom(value),
         })
     }
     c.JSON(200, gin.H{
           "message":"Query Success",
     })
}

	 

