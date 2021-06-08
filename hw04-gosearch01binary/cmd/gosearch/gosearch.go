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

	var ind = index.Storage(make(map[string][]int))
	ind.Store(docs)

	if *token != "" {
		fmt.Println("Search results:")
		ids, ok := ind[strings.ToLower(*token)]
		if !ok {
			fmt.Println("Not found")
			return
		}
		for _, id := range ids {
			x := sort.Search(len(docs), func(i int) bool {
				return docs[i].ID >= id
			})
			fmt.Println(docs[x].URL, docs[x].Title)
		}
	}
}
