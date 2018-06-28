package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("you cuo wu!", err)
	}
}
func main() {
	opening, err := goquery.NewDocument("https://baidu.com/s?wd=微信")
	checkErr(err)

	ele := opening.Find("div").Find(".result").Find("h3")
	
	defer ele.Each(func(i int, content *goquery.Selection) {
		fmt.Println()
		a := content.Text()
		a = strings.Replace(a, " ", "", -1)
		a = strings.Replace(a, "\n", "", -1)
		fmt.Print(a, "    ")
		b, _ := content.Find("a").Attr("href")
		fmt.Println(b)
	})
	ele1 := opening.Find("div").Find(".result-op").Find("h3")
	defer ele1.Each(func(i int, content *goquery.Selection) {
		fmt.Println()
		a := content.Text()
		a = strings.Replace(a, " ", "", -1)
		a = strings.Replace(a, "\n", "", -1)
		fmt.Print(a, "    ")
		b, _ := content.Find("a").Attr("href")
		fmt.Println(b)
	})
	defer fmt.Println()
}
