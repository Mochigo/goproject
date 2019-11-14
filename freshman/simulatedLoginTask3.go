package main

import(
      "fmt"
      "net/http"
      "io/ioutil"
      "net/url"
      "strings"

      "github.com/howeyc/gopass"
      "github.com/tidwall/gjson"
)

func main() {
     requestUrl := "https://accounts.douban.com/j/mobile/login/basic"

     var name string
     fmt.Printf("豆瓣账号：")
     _, _ = fmt.Scanln(&name)
     fmt.Printf("输入密码：")
     password, _ := gopass.GetPasswdMasked()

     data := url.Values{}
     data.Set("name", name)
     data.Set("password", string(password))
     data.Set("remember", "false")
     data.Set("ck", "")
     data.Set("ticket", "")

     payload := strings.NewReader(data.Encode())

     req, err := http.NewRequest("POST", requestUrl, payload)
     if err != nil {
	panic(err)
	return
     }

     req.Header.Add("Accept", "application/json")
     req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
     req.Header.Add("Origin", "https://accounts.douban.com")
     req.Header.Add("Referer", "https://accounts.douban.com/passport/login_popup?login_source=anony")
     req.Header.Add("Sec-Fetch-Mode", "cors")
     req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")
     req.Header.Add("X-Requested-With", "XMLHttpRequest")

     res, err := http.DefaultClient.Do(req)
     if err != nil {
	panic(err)
	return
     }

     defer res.Body.Close()
     body, err := ioutil.ReadAll(res.Body)
     if err != nil {
	panic(err)
	return
     }
     fmt.Println(res.Status)

     result := gjson.Get(string(body), "message")
     fmt.Println(result)
     fmt.Println(gjson.Get(string(body), "description"))

     fmt.Println(string(body))
}




