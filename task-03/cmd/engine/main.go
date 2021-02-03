package main

import (
	"dersnek/task-3/pkg/spider"
	"fmt"
	"strings"
)

// Scanner - interface for scanners/crawlers
type Scanner interface {
	Scan(url string) (data map[string]string, err error)
}

func main() {
	urls := []string{"https://katata.games", "https://go.dev"}
	store := make(map[string]string)
	scanner := spider.New(2)
	store, err := scan(scanner, store, urls)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		q := ""

		fmt.Printf("\n------------------------------------\n")
		fmt.Printf("Search query: ")
		fmt.Scanln(&q)
		fmt.Printf("------------------------------------\n")

		if q == "" {
			continue
		}

		q = strings.ToLower(q)
		res := make(map[string]string)

		for k, v := range store {
			if strings.Contains(strings.ToLower(k), q) || strings.Contains(strings.ToLower(v), q) {
				res[k] = v
			}
		}

		fmt.Printf("\n%v results\n", len(res))
		for k, v := range res {
			fmt.Println()
			fmt.Println(v)
			fmt.Println(k)
		}
	}
}

func scan(scanner Scanner, store map[string]string, urls []string) (map[string]string, error) {
	for i := 0; i < len(urls); i++ {
		fmt.Printf("Scanning %s...", urls[i])

		data, err := scanner.Scan(urls[i])
		if err != nil {
			return nil, err
		}

		for k, v := range data {
			store[k] = v
		}

		fmt.Printf(" done!\n")
	}

	return store, nil
}
