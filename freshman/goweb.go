package main

import(
      "fmt"
      "encoding/json"
      "io/ioutil"
      "net/http"
)
//记录用户信息
type User struct {
     Id        uint
     Name      string
     Password  string
}
//记录状态
type Status struct {
     State    bool
     Detail   string
}
//利用数组记录各个用户
var userArr = make([]User, 0)
//给用户发放id（排序）
var userId uint = 1

var status = Status{false, ""}

func Existed(user User) bool{
     for _, value := range userArr {
         if value.Name == user.Name {
            return true
	 }
     }
     return false
}

func Verify(user User) bool {
     for _, value := range userArr {
	 if value.Name == user.Name && value.Password == user.Password {
	    return true 
         }
     }
     return false
}

func Register(userInfo []byte) {
     var user User
     json.Unmarshal(userInfo, &user)
     if Existed(user) {
	status = Status{false, "用户名已存在"}
	return
     }
     user.Id = userId
     userId += 1
     userArr = append(userArr, user)
     status = Status{true, "注册成功"}
}

func LoginIn(userInfo []byte) {
     var user User
     json.Unmarshal(userInfo, &user)
     if !Existed(user) {
	 status = Status{false, "用户名不存在"}
	 return 
     }
     status = Status{true, "登陆成功" }
}

func Reset(userInfo []byte) {
     var user User
     json.Unmarshal(userInfo, &user)
     for _, tmp := range userArr {
	 if user.Name == tmp.Name {
            tmp.Password = user.Password
	    status = Status{true, "修改成功"}
	    fmt.Println("修改成功！")
	    return
         }
     }
     status = Status{false, "账号不存在。"}
}

func Feedbook(finfo Status) []byte {
     s, _ := json.Marshal(finfo)
     return s
}

func register(res http.ResponseWriter, req *http.Request) {
     if req.Method == "POST" {
	s, _ := ioutil.ReadAll(req.Body)
	Register(s)
	res.Write(Feedbook(status))
     } else {
	res.Write([]byte("\"false\":\"只支持POST方式\"}"))
     }
}

func login(res http.ResponseWriter, req *http.Request) {
     if req.Method == "POST" {
	s, _ := ioutil.ReadAll(req.Body)
	LoginIn(s)
	res.Write(Feedbook(status))
     } else {
	res.Write([]byte("{\"false\":\"只支持POST方法\"}"))
     }
}

func reset(res http.ResponseWriter, req *http.Request) {
     if req.Method == "POST" {
	s, _ := ioutil.ReadAll(req.Body)
        Reset(s)
	res.Write(Feedbook(status))
     } else {
	res.Write([]byte("{\"false\":\"只支持POST方法\"}"))
     }
}

func main() {
     http.HandleFunc("/login", login)
     http.HandleFunc("/register", register)
     http.HandleFunc("/reset", reset)
     if err := http.ListenAndServe(":9090",nil); err != nil{
	fmt.Println("监听失败")
     }
}





















