package memstore

import (
	"go-search/pkg/store"
)

// Service is a file store service.
type Service struct {
	store.Interface
	Memory []*store.Doc
}

// New creates and returns a new memstore service.
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

// FindInMemory - true implementation of Find method for memstore,
// returns a docs from Memory array with a given id
// Used only for benchmarking
func (s *Service) FindInMemory(id uint) *store.Doc {
	if int(id) < len(s.Memory) {
		return s.Memory[id]
	}

	return nil
}

// Add appends doc to the Memory array of docs in the memstore
func (s *Service) Add(doc *store.Doc, _ bool) uint {
	s.Memory = append(s.Memory, doc)
	return uint(len(s.Memory) - 1)
}
