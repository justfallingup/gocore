package storage

import (
	"encoding/json"
	"github.com/justfallingup/gocore/hw06/pkg/crawler"
	"io"
	"os"
	"sort"
)

type JsonStore struct {
	F  *os.File
}

func Store(w io.Writer, docs []crawler.Document) error {
	docs = sortDocs(docs)
	b, err := json.Marshal(docs)
	if err != nil {
		return err
	}
	n, err := w.Write(b)
	if n != len(b) {
		return io.ErrShortWrite
	}
	return err
}

func Get(r io.Reader) ([]crawler.Document, error) {
	var b []byte
	var buf = make([]byte, 42)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		buf = buf[:n]
		b = append(b, buf...)
	}
	var d []crawler.Document
	err := json.Unmarshal(b, &d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func sortDocs(scanDocs []crawler.Document) []crawler.Document {
	var docs []crawler.Document
	for i, d := range scanDocs {
		d.ID = i
		docs = append(docs, d)
	}
	sort.Slice(docs, func(i, j int) bool {
		return docs[i].ID < docs[j].ID
	})
	return docs
}
