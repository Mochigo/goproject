package register

import (
"github.com/gin-gonic/gin"
"github.com/MitsuhaOma/goproject/goGin/model"
)

type RegisterPayload strust {

     Username     string `json:"name"`
     Password string `json:"password"`

}

func RegisterStudent(c *gin.Context) {
     var data RegisterPayload
     //BindJSON会将c中内容以json的格式存储在data中
     if err := c.BindJSON(&data); err != nil {
	c.JOSN(400, gin.H{
	      "messsage": "Bad Request!",
        })
	return
     }
     if model.CheckUserByUsername(data.Name) {
	c.JSON(401, gin.H{
	      "message": "User Already Existed!",
        })
	return
     }
     model.CreateUser(data.Name, data.Password)
     c.JSON(200, gin.H{
           "message": "Register Successful!",
     })
     return
}
