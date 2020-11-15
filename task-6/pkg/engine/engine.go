package engine

import (
	"dersnek/task-6/pkg/crawler"
	"dersnek/task-6/pkg/crawler/spider"
	"dersnek/task-6/pkg/index"
	"dersnek/task-6/pkg/store"
	"dersnek/task-6/pkg/store/file"
	"fmt"
	"strings"
)

var urls = []string{"https://katata.games", "https://go.dev"}

// Service - search engine service
type Service struct {
	i *index.Service
}

// New - creates and returns a new engine service
func New() *Service {
	return &Service{}
}

// Start - fires up the search engine.
// You have to run Start() first before searching anything
func (s *Service) Start() error {
	store := file.New()

	savedDocs, savedIndex, err := store.Read()
	if err != nil {
		return err
	}

	s.i = index.New()
	scanner := spider.New()

	if len(savedDocs) == 0 || len(savedIndex) == 0 {
		fmt.Println("Cache not found.")
		fmt.Println("Running initial scan, this might take a while...")

		scanRes, err := scan(scanner, urls, true)
		if err != nil {
			return err
		}

		docs := s.i.Build(scanRes)
		store.Write(docs, s.i.Data)
	} else {
		fmt.Println("Cache found!")

		fmt.Println("Restoring from cache...")
		s.i.Restore(savedDocs, savedIndex)

		fmt.Println("Starting a background scan...")
		go s.bgScan(scanner, s.i, store, urls)
	}

	return nil
}

// Find - finds documents (webpages) containing words present in the query string
func (s *Service) Find(q string) map[string]string {
	res := make(map[string]string)
	qWords := strings.Split(strip(q), " ")

	for _, w := range qWords {
		wRes := findByWord(w, s.i)

		for _, d := range wRes {
			if res[d.URL] != d.Title {
				res[d.URL] = d.Title
			}
		}
	}

	return res
}

func findByWord(s string, i *index.Service) []store.Doc {
	res := []store.Doc{}

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

func (s *Service) bgScan(scanner crawler.Interface, indexer *index.Service, store *file.Service, urls []string) {
	scanRes, err := scan(scanner, urls, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	s.i = index.New()
	docs := s.i.Build(scanRes)
	store.Write(docs, s.i.Data)
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
