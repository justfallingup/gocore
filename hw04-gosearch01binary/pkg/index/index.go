package index

import (
	"github.com/justfallingup/gocore/hw04-gosearch01binary/pkg/crawler"
	"sort"
	"strings"
)

type Storage map[string][]int

func New() Storage {
	return make(map[string][]int)
}

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

func (s Storage) Get(word string, docs []crawler.Document) []crawler.Document {
	ids, ok := s[word]
	if !ok {
		return nil
	}
	var res []crawler.Document
	for _, id := range ids {
		x := sort.Search(len(docs), func(i int) bool {
			return docs[i].ID >= id
		})
		res = append(res, docs[x])
	}
	return res
}

func stored(s []int, n int) bool {
	for _, i := range s {
		if n == i {
			return true
		}
	}
	return false
}

