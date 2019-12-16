package main

import (
       "fmt"
       "github.com/MitsuhaOma/goproject/goGin/model"
       "github.com/MitsuhaOma/goproject/goGin/router"
)

func main() {
     model.Init()
     router.InitRouter()
     router.Router.Run()
     fmt.Println("Running... Successful!")
}
