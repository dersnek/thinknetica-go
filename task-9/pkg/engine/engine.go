package engine

import (
	"fmt"
	"go-search/pkg/crawler"
	"go-search/pkg/index"
	"go-search/pkg/index/hash"
	"go-search/pkg/store"
	"strings"
)

// Service - search engine service
type Service struct {
	ind index.Interface
}

// New - creates and returns a new engine service
func New(ind index.Interface) *Service {
	return &Service{
		ind: ind,
	}
}

// Start - fires up the search engine.
// You have to run Start() first before searching anything
func (s *Service) Start(store store.Interface, scanner crawler.Interface, urls []string) error {
	savedDocs, savedIndex, err := store.Read()
	if err != nil {
		return err
	}

	if len(savedDocs) == 0 || len(savedIndex.(index.Data)) == 0 {
		scanRes, err := scan(scanner, urls, true)
		if err != nil {
			return err
		}
		docs := s.ind.Build(scanRes, store)
		store.Write(docs, s.ind.Index())
	} else if len(savedDocs) > 0 && len(savedIndex.(index.Data)) > 0 {
		s.ind.Restore(savedDocs, savedIndex.(index.Data), store)
		go s.bgScan(scanner, s.ind, store, urls)
	}

	return nil
}

// Find - finds documents (webpages) containing words present in the query string
func (s *Service) Find(q string, st store.Interface) []store.Doc {
	res := []store.Doc{}
	qWords := strings.Split(strip(q), " ")

	for _, w := range qWords {
		wRes := findByWord(w, s.ind, st)

		for _, wPage := range wRes {
			dup := false
			for _, page := range res {
				if page.Title == wPage.Title {
					dup = true
				}
			}
			if !dup {
				res = append(res, wPage)
			}
		}
	}

	return res
}

func findByWord(s string, ind index.Interface, st store.Interface) []store.Doc {
	res := []store.Doc{}

	ids, found := ind.Find(s)
	if !found {
		return res
	}

	for _, id := range ids {
		d := st.Find(id)

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

func (s *Service) bgScan(scanner crawler.Interface, indexer index.Interface, store store.Interface, urls []string) {
	scanRes, err := scan(scanner, urls, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	s.ind = hash.New()
	docs := s.ind.Build(scanRes, store)
	store.Write(docs, s.ind.Index())
}

func scan(scanner crawler.Interface, urls []string, verbose bool) ([]store.Doc, error) {
	res := []store.Doc{}

	for i := 0; i < len(urls); i++ {
		if verbose {
			fmt.Printf("Scanning %s...", urls[i])
		}

		data, err := scanner.Scan(urls[i], 2)
		if err != nil {
			return nil, err
		}

		for _, d := range data {
			res = append(res, d)
		}

		if verbose {
			fmt.Printf(" done!\n")
		}
	}

	return res, nil
}
