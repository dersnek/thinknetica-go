package engine

import (
	"dersnek/task-5/pkg/index"
	"dersnek/task-5/pkg/w"
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

func findByWord(s string, i *index.Service) []w.Doc {
	res := []w.Doc{}

	ids, found := i.Data[s]
	if !found {
		return res
	}

	for _, id := range ids {
		d := i.Store.Find(id)

		if d != nil {
			res = append(res, *d)
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
