// Package hash indexes, stores and retrieves webpages
package hash

import (
	"dersnek/task-8/pkg/index"
	"dersnek/task-8/pkg/store"
	"strings"
)

// Service - search index service
type Service struct {
	index index.Data
}

// New - creates and returns a new index service
func New() *Service {
	return &Service{
		index: make(index.Data),
	}
}

type wordIDPair struct {
	word string
	id   uint
}

var stopWords = []string{
	"a", "and", "around", "every", "for", "from", "in",
	"is", "it", "not", "on", "one", "the", "to", "under"}

// Index returns data, which is the actual index
func (s *Service) Index() index.Data {
	return s.index
}

// Build builds an index for an array of store.Docs
// Returns array of same docs, but with assigned IDs
func (s *Service) Build(docs []store.Doc, store store.Interface) []store.Doc {
	allWordIDPairs := []wordIDPair{}

	// Extract all words from all docs into an array
	for i := 0; i < len(docs); i++ {
		// Store doc
		id := store.Add(&docs[i], true)
		docs[i].ID = id

		docWords := strings.Split(strip(docs[i].Title), " ")

		for _, w := range docWords {
			widp := wordIDPair{word: w, id: id}
			allWordIDPairs = append(allWordIDPairs, widp)
		}
	}

	// Convert all words to lowercase
	for i, wdp := range allWordIDPairs {
		allWordIDPairs[i].word = strings.ToLower(wdp.word)
	}

	// Filter out stop words
	filtWordIDPairs := []wordIDPair{}

	for _, widp := range allWordIDPairs {
		if !contains(stopWords, widp.word) {
			filtWordIDPairs = append(filtWordIDPairs, widp)
		}
	}

	for _, widp := range filtWordIDPairs {
		s.index[widp.word] = append(s.index[widp.word], widp.id)
	}

	return docs
}

// Restore - assigns index.Data (index) to the service and rebuils the binary tree
func (s *Service) Restore(docs []store.Doc, data index.Data, store store.Interface) {
	s.index = data
	for i := 0; i < len(docs); i++ {
		store.Add(&docs[i], false)
	}
}

// Find returns an array of ids of documents which include the word
func (s *Service) Find(word string) ([]uint, bool) {
	ids, found := s.index[word]
	return ids, found
}

func strip(s string) string {
	var result strings.Builder
	result.Grow(len(s))
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}
	return result.String()
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
