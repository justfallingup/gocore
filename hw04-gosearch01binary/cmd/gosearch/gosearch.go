package main

import (
	"flag"
	"fmt"
	"github.com/justfallingup/gocore/hw04-gosearch01binary/pkg/crawler"
	"github.com/justfallingup/gocore/hw04-gosearch01binary/pkg/crawler/spider"
	"github.com/justfallingup/gocore/hw04-gosearch01binary/pkg/index"
	"log"
	"sort"
	"strings"
)

func main()  {
	token := flag.String("s", "", "a word you're searching for")
	flag.Parse()
	if *token == "" {
		flag.PrintDefaults()
		return
	}

	urls := []string{
		"https://go.dev",
		"https://golang.org",
	}
	s := spider.New()
	const depth = 2
	var scanDocs []crawler.Document
	for _, u := range urls {
		res, err := s.Scan(u, depth)
		if err != nil {
			log.Println("scanner error")
			continue
		}
		scanDocs = append(scanDocs, res...)
	}

	var docs []crawler.Document
	for i, d := range scanDocs {
		d.ID = i
		docs = append(docs, d)
	}
	sort.Slice(docs, func(i, j int) bool {
		return docs[i].ID < docs[j].ID
	})

	ind := index.New()
	ind.Store(docs)

	fmt.Println("Search results:")
	res := ind.Get(strings.ToLower(*token), docs)
	if res == nil {
		fmt.Println("Not found")
		return
	}
	for _, i := range res {
		fmt.Println(i.URL, i.Title)
	}
}
