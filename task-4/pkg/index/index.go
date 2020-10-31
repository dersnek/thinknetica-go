// Package index indexes, stores and retrieves webpages
package index

import (
	"dersnek/task-4/pkg/crawler"
	"sort"
	"strings"
)

// Service - search index service
type Service struct {
	Store  []crawler.Document
	Data   Data
	LastID uint
}

// New - creates and returns a new index service
func New() *Service {
	return &Service{
		Data:   make(Data),
		LastID: 0,
	}
}

// Data contains a hashmap with words as keys and as values -
// an array of IDs of documents (webpages) containing given word.
type Data map[string][]uint

type wordIDPair struct {
	word string
	id   uint
}

var stopWords = []string{
	"a", "and", "around", "every", "for", "from", "in",
	"is", "it", "not", "on", "one", "the", "to", "under"}

// Process stores array of crawler.Documents
// and build an inverted index for them
func (s *Service) Process(docs []crawler.Document) {
	allWordIDPairs := []wordIDPair{}

	// Extract all words from all docs into an array
	for i, d := range docs {
		gID := s.LastID + uint(i)

		// Store doc
		d.ID = gID
		s.Store = append(s.Store, d)

		docWords := strings.Split(strip(d.Title), " ")

		for _, w := range docWords {
			widp := wordIDPair{word: w, id: gID}
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
		s.Data[widp.word] = append(s.Data[widp.word], widp.id)
	}

	sort.Slice(s.Store, func(i, j int) bool { return s.Store[i].ID < s.Store[j].ID })
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
