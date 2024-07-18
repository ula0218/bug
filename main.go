package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// request(GET)
	req, err := http.NewRequest("GET", "https://www.ptt.cc/bbs/Stock/index.html", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("User-Agent", "MyCustomCrawler/1.0")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	//  HTML 文檔中 class 屬性為 title 的 <div> 元素
	doc.Find("div.title").Each(func(i int, item *goquery.Selection) {
		title := item.Text()
		fmt.Printf("%s\n", title)
	})
}
