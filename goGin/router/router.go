package router

import (
	"github.com/gin-gonic/gin"
	"github.com/MitsuhaOma/goproject/goGin/handler"
)

var Router *gin.Engine

func InitRouter() {
     Router = gin.Default()
     Router.POST("/student/create",handler.CreateStudent)
  

     return
}
