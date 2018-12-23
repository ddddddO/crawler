package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	gq "github.com/PuerkitoBio/goquery"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "health OK!")
	})

	r.POST("/", PostHandler)

	r.Run(":4567")
}

func PostHandler(c *gin.Context) {
	c.Request.ParseForm() // リクエストフォームをパース
/*
	fmt.Println("---post request [targetURL]---")
	var targetURL string
	targetURL = c.Request.Form["targetURL"][0]
	fmt.Println(targetURL)
	fmt.Println()
*/
	fmt.Println("---post request [tURL]---")
	tURL := c.Request.Form["tURL"][0]
	fmt.Println(tURL)
	fmt.Println()

	// クライアントから選択したクロール対象サイトごとに、クロール処理を分岐させる


	fmt.Println("---start crawl---")
	// クロール対象urlを記事一覧サイトから取得
	listDoc, err := gq.NewDocument("http://blog.livedoor.jp/nwknews/")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	utMap := map[string]string{}
	listSlc := listDoc.Find("div > div > div.article-header > div.article-title-outer > h2")
	listSlc.Each(func(index int, s *gq.Selection) {
		var u string
		u = s.Find("a").AttrOr("href", "")

		var t string
		t = s.Text()

		utMap[u] = t
	})

	fmt.Println("---ListURLs---")
	for u, t := range utMap {
		fmt.Printf("URL: %s\n", u)
		fmt.Printf("TITLE: %s\n", t)
		fmt.Println()
	}

	fmt.Println("---end crawl---")
	c.String(200, tURL)
}
