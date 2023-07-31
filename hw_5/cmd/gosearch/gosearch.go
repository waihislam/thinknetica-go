package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"thinknetica-go/hw_5/pkg/crawler"
	"thinknetica-go/hw_5/pkg/crawler/spider"
	ind "thinknetica-go/hw_5/pkg/index"
)

var url1 = "https://go.dev"
var url2 = "https://www.practical-go-lessons.com/"

func main() {
	kw := flag.String("s", "Go", "keyword")
	flag.Parse()
	log.Printf("Trying to find keyword: %s\n", *kw)

	var index *ind.Index
	path := "hw_5/attach/index.json"
	if checkPath(path) {
		readFile(path, index)
	}

	if index.IsEmpty() {
		res, err := writeFile(path, *kw)
		if err != nil {
			return
		}
		// print result
		if len(res) != 0 {
			fmt.Printf("Found %v urls\n", len(res))
			for _, doc := range res {
				fmt.Printf("URL: %v, Title: %v\n", doc.URL, doc.Title)
			}
		} else {
			log.Println("Nothing found")
		}
	}
}

func readFile(path string, index *ind.Index) {
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	fIndex, err := ind.GetData(f)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(fIndex, &index)
	if err != nil {
		log.Println(err)
		return
	}
}

func writeFile(path string, kw string) ([]*crawler.Document, error) {
	urls := [2]string{url1, url2}
	// scanning urls
	var docs []crawler.Document
	s := spider.New()

	for _, url := range urls {
		doc, err := s.Scan(url, 2)
		if err != nil {
			log.Println(url, err)
			continue
		}
		docs = append(docs, doc...)

	}

	// binary search by index
	index := ind.New()
	index.AddDocuments(docs)
	sort.Sort(index)

	f, err := os.Create(path)
	if err != nil {
		log.Println("File not created", err)
		return nil, err
	}
	defer f.Close()

	err = index.WriteJson(f)
	if err != nil {
		log.Println("Can't write to file", err)
		return nil, err
	}

	IDs := index.Search(kw)
	res := index.GetDocsByID(IDs)
	return res, nil
}

func checkPath(path string) bool {
	found := false

	if _, err := os.Stat(path); err == nil {
		found = true
	} else {
		log.Println(err)
	}
	return found
}
