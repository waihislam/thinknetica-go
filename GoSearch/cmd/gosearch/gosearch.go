package main

import (
	"flag"
	"fmt"
	"strings"
	"thinknetica-go/GoSearch/pkg/crawler"
	"thinknetica-go/GoSearch/pkg/crawler/spider"
)

var url1 = "https://go.dev"
var url2 = "https://www.practical-go-lessons.com/"

func main() {
	urls := [2]string{url1, url2}
	kw := flag.String("s", "", "keyword")
	flag.Parse()
	fmt.Printf("Trying to find keyword: %s\n", *kw)

	// scanning urls
	var docs []crawler.Document
	s := spider.New()

	for _, url := range urls {
		doc, err := s.Scan(url, 2)
		if err != nil {
			fmt.Println(url, err)
			continue
		}
		docs = append(docs, doc...)

	}

	// search the word
	var res []string
	for _, doc := range docs {
		if strings.Contains(strings.ToLower(doc.Title), strings.ToLower(*kw)) {
			res = append(res, doc.URL)
		}
	}

	// print result
	if len(res) != 0 {
		fmt.Printf("Found %v urls\n", len(res))
		fmt.Println(res)
	} else {
		fmt.Println("Nothing found")
	}
}
