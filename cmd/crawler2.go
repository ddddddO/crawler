package main

import (
	"fmt"
	"os"

	rf "github.com/ddddddO/crawler/readfile"

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
	fmt.Println("---post request---")
	var targetURL string
	targetURL = c.Request.Form["targetURL"][0]
	fmt.Println(targetURL)
	fmt.Println()

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


	// 個別クロール対象urlをファイルから取得
	tUrls, err := rf.ReadFile()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("---哲学ニュース title---")
	for _, tUrl := range tUrls {
		doc, err := gq.NewDocument(tUrl)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		slc := doc.Find("#main > div > div > div.article-outer.hentry > div > div > div.article-header > div.article-title-outer > h2")

		fmt.Printf("Title: %s\n", slc.Text())
		fmt.Printf("URL  : %s\n", tUrl)
		fmt.Println()
	}

	fmt.Println("---end crawl---")
}
