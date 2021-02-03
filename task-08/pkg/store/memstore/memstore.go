package memstore

import (
	"dersnek/task-8/pkg/store"
)

// Service is a file store service.
type Service struct {
	store.Interface
	Memory []*store.Doc
}

// New creates and returns a new file store service.
func New() *Service {
	s := Service{}
	return &s
}

// Find - find return a web document, a fake search result
func (s *Service) Find(id uint) *store.Doc {
	var doc store.Doc
	switch id {
	case 0:
		doc = store.Doc{URL: "https://katata.games", Title: "Katata Games"}
	case 1:
		doc = store.Doc{URL: "https://go.dev", Title: "Go"}
	}
	return &doc
}

// Add appends doc to the Memory array of docs in the memstore
func (s *Service) Add(doc *store.Doc, _ bool) uint {
	s.Memory = append(s.Memory, doc)
	return 0
}
