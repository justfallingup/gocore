package index

import (
	"github.com/justfallingup/gocore/hw04-gosearch01binary/pkg/crawler"
	"strings"
)

type Storage map[string][]int

func (s Storage) Store(docs []crawler.Document)  {
	for _, doc := range docs {
		words := strings.Fields(doc.Title)
		for _, word := range words {
			word = strings.ToLower(word)
			if i, ok := s[word]; ok && stored(i, doc.ID){
				continue
			}
			s[word] = append(s[word], doc.ID)
		}
	}
}

func stored(s []int, n int) bool {
	for _, i := range s {
		if n == i {
			return true
		}
	}
	return false
}

