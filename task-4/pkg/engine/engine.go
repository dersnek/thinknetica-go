package engine

import (
	"dersnek/task-4/pkg/crawler"
	"dersnek/task-4/pkg/index"
	"sort"
	"strings"
)

// Service - search engine service
type Service struct{}

// New - creates and returns a new engine service
func New() *Service {
	return &Service{}
}

// Find - finds documents (webpages) containing words present in the query string
func (s *Service) Find(q string, i *index.Service) map[string]string {
	res := make(map[string]string)
	qWords := strings.Split(strip(q), " ")

	for _, w := range qWords {
		wRes := findByWord(w, i)

		for _, d := range wRes {
			if res[d.URL] != d.Title {
				res[d.URL] = d.Title
			}
		}
	}

	return res
}

func findByWord(w string, s *index.Service) []crawler.Document {
	res := []crawler.Document{}

	ids, found := s.Data[w]
	if !found {
		return res
	}

	for _, id := range ids {
		i := sort.Search(len(s.Store), func(i int) bool { return (s.Store[i]).ID >= id })
		if i < len(s.Store) && s.Store[i].ID == id {
			res = append(res, s.Store[i])
		}
	}

	return res
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
