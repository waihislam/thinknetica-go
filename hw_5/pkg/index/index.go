package index

import (
	"encoding/json"
	"io"
	"sort"
	"strings"
	"thinknetica-go/hw_5/pkg/crawler"
)

type Index struct {
	Docs   []*crawler.Document `json:"documents"`
	IndMap map[string][]int    `json:"indMap"`
}

func (idx *Index) Len() int {
	return len(idx.Docs)
}

func (idx *Index) Less(i, j int) bool {
	return idx.Docs[i].ID < idx.Docs[j].ID
}

func (idx *Index) Swap(i, j int) {
	idx.Docs[i], idx.Docs[j] = idx.Docs[j], idx.Docs[i]
}

func New() *Index {
	return &Index{
		Docs:   []*crawler.Document{},
		IndMap: make(map[string][]int),
	}
}

func (idx *Index) AddDocuments(docs []crawler.Document) {
	idx.Docs = make([]*crawler.Document, len(docs))
	for i, doc := range docs {
		idx.Docs[i] = &docs[i]
		words := strings.Fields(doc.Title)
		for _, word := range words {
			word = strings.ToLower(word)
			if !findElement(idx.IndMap[word], doc.ID) {
				idx.IndMap[word] = append(idx.IndMap[word], doc.ID)
			}
		}
	}
}

func (idx *Index) Search(word string) []int {
	word = strings.ToLower(word)
	return idx.IndMap[word]
}

func (idx *Index) GetDocsByID(ids []int) []*crawler.Document {
	var docs []*crawler.Document
	for _, id := range ids {
		i := sort.Search(len(idx.Docs), func(i int) bool {
			return idx.Docs[i].ID >= id
		})
		docs = append(docs, idx.Docs[i])
	}
	return docs
}

func findElement(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func GetData(r io.Reader) ([]byte, error) {
	fileData, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return fileData, err
}

func (idx *Index) WriteJson(w io.Writer) error {
	jsonData, err := json.Marshal(idx)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonData)
	return err
}

func (idx *Index) IsEmpty() bool {
	return idx == nil || (len(idx.Docs) == 0 && len(idx.IndMap) == 0)
}
