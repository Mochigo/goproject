package handler

import(
      "github.com/gin-gonic/gin"
      "github.com/MitsuhaOma/goproject/goGin/model"
)


func CreatStudent(c *gin.Context) {
     var tmpStudent students
     if err := c.BindJSON(&tmpStudent); err != nil {
	c.JSON(400,gin.H{
	      "message":"Bad Request!",
        })
	return
     }
										for _,v := range tmpStudent.Student {
	 c.JSON(200,gin.H{
		"message":model.CreatStudent(v),
	 })
     }
     c.JSON(200,gin.H{
	    "message":"Creat Success",
     })
}
