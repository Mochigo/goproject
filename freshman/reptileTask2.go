package main

import(
      "fmt"
      "strings"
      "github.com/MitsuhaOma/goquery"
      "net/http"
      "io/ioutil"
)

func main() {
	num := 1
	requestUrl := "https://studygolang.com/topics"
	rp, err := http.Get(requestUrl)
	if err != nil {
	   panic(err)
        }
	body, err := ioutil.ReadAll(rp.Body)
	if err != nil {
           panic(err)
	}
	content := string(body)
	defer rp.Body.Close()

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
	   panic(err)
        }

	dom.Find("div.title").Each(func(i int, selection *goquery.Selection) {
	    selection.Find("a[href][title]").Each(func(i int, title *goquery.Selection) {
		    fmt.Printf("第%2d个文档的标题:", num)
		    fmt.Println(title.Text())
		    num++
	    })
	})
}
