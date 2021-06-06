package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/justfallingup/gocore/hw03/pkg/crawler"
	"github.com/justfallingup/gocore/hw03/pkg/crawler/spider"
)

func main()  {
	userSearch := flag.String("s", "", "-s 'word from title of web-page'")
	flag.Parse()
	userString := strings.ToLower(*userSearch)

	urls := []string{
		"https://go.dev/",
		"https://golang.org/",
	}

	var server crawler.Interface = spider.New()

	var docs []crawler.Document
	for _, u := range urls {
		d, err := server.Scan(u, 2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		docs = append(docs, d...)
	}

	for _, v := range docs {
		if userString != "" {
			if strings.Contains(strings.ToLower(v.Title), userString) {
				fmt.Println(v.URL, v.Title)
			}
			continue
		}
		fmt.Println(v.URL, v.Title)
	}
}
