package handler

import (
       "github.com/gin-gonic/gin"
       "github.com/MitsuhaOma/goproject/goGin/model"
       "log"
)

type students struct {
	Student []string `json:"student"`
}

func CreateStudent(c *gin.Context) {
     var tmpStudent students
     if err := c.BindJSON(&tmpStudent); err != nil {
	c.JSON(400, gin.H{
	       "message":"Bad Request!",
	})      
	return
     }
     var student
     for _, value := range tmpStudent.Student {
	 c.JSON(200, gin.H{
		 "message": model.CreateStudent(value),
         })
     }
     c.JSON(200, gin.H{
	    "message": "Create Success!",
     })
}
