package main

import (
	"flag"
	"fmt"
	"github.com/justfallingup/gocore/hw03-gosearch01/pkg/crawler"
	"github.com/justfallingup/gocore/hw03-gosearch01/pkg/crawler/spider"
	"log"
	"strings"
)

func main()  {
	token := flag.String("s", "", "a word you're searching for")
	flag.Parse()
	urls := []string{
		"https://go.dev",
		"https://golang.org",
	}
	if *token != "" {
		s := spider.New()
		const depth = 2
		var docs []crawler.Document
		for _, u := range urls {
			res, err := s.Scan(u, depth)
			if err != nil {
				log.Println("scanner error")
				continue
			}
			docs = append(docs, res...)
		}
		fmt.Println("Search results:")
		for _, d := range docs {
			if strings.Contains(strings.ToLower(d.Title), strings.ToLower(*token)) {
				fmt.Println(d.URL, d.Title)
			}
		}
	}
}
