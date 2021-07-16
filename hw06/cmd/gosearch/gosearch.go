package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/justfallingup/gocore/hw06/pkg/crawler"
	"github.com/justfallingup/gocore/hw06/pkg/crawler/spider"
	"github.com/justfallingup/gocore/hw06/pkg/index"
	"github.com/justfallingup/gocore/hw06/pkg/storage/jsonfile"
	"log"
	"os"
	"strings"
)

func main()  {
	const depth = 2
	urls := []string{
		"https://go.dev",
		"https://golang.org",
	}
	var docs []crawler.Document
	db := storage.JsonStore{}
	ind := index.New()
	s := spider.New()
	token := flag.String("s", "", "a word you're searching for")
	filename := "./" + *flag.String("f", "docs.json", "filename.json")

	flag.Parse()
	if *token == "" {
		flag.PrintDefaults()
		return
	}
	err := os.Chdir(os.Getenv("GOPATH") + "/src/github.com/justfallingup/gocore/hw06/pkg/storage/jsonfile")
	if err != nil {
		log.Println("change dir error", err)
		return
	}
	_, errFile := os.Stat(filename)
	if errFile != nil && !errors.Is(errFile, os.ErrNotExist) {
		log.Println("check file error", errFile)
		return
	}
	if errors.Is(errFile, os.ErrNotExist) {
		fmt.Println("Scanning urls...")
		for _, u := range urls {
			res, err := s.Scan(u, depth)
			if err != nil {
				log.Println("scanner error", err)
				continue
			}
			docs = append(docs, res...)
		}

		f, err := os.Create(filename)
		if err != nil {
			log.Println("file creation error", err)
			return
		}
		defer f.Close()
		db.F = f
		err = storage.Store(db.F, docs)
		if err != nil {
			log.Println("storage store error", err)
			return
		}
	}
	f, err := os.Open(filename)
	if err != nil {
		log.Println("file open error", err)
	}
	defer f.Close()
	db.F = f
	docs, err = storage.Get(db.F)
	if err != nil {
		log.Println("get storage error", err)
	}

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
