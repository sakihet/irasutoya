package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	query := os.Args[1]
	doc, err := goquery.NewDocument("http://www.irasutoya.com/search?q=" + query)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".boxim").Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Find("a").Attr("href")
		doc, _ := goquery.NewDocument(href)
		h2 := doc.Find(".title h2").Text()
		title := strings.Replace(h2, "\n", "", -1)
		href2, _ := doc.Find(".entry .separator a").Attr("href")
		fmt.Println(title)
		fmt.Println(href2)
	})
}
