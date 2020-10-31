package main

import (
	"bufio"
	"dersnek/task-4/pkg/crawler"
	"dersnek/task-4/pkg/crawler/spider"
	"dersnek/task-4/pkg/engine"
	"dersnek/task-4/pkg/index"
	"fmt"
	"os"
)

func main() {
	urls := []string{"https://katata.games", "https://go.dev"}
	indexer := index.New()
	scanner := spider.New()
	engine := engine.New()

	scanRes, err := scan(scanner, urls)
	if err != nil {
		fmt.Println(err)
		return
	}

	indexer.Process(scanRes)

	for {
		q := ""

		fmt.Printf("\n------------------------------------\n")
		fmt.Printf("Search query: ")

		bscanner := bufio.NewScanner(os.Stdin)
		for bscanner.Scan() {
			q = bscanner.Text()
			break
		}

		fmt.Printf("------------------------------------\n")

		if q == "" {
			continue
		}

		res := engine.Find(q, indexer)

		printResults(res)
	}
}

func scan(scanner crawler.Interface, urls []string) ([]crawler.Document, error) {
	res := []crawler.Document{}

	for i := 0; i < len(urls); i++ {
		fmt.Printf("Scanning %s...", urls[i])

		data, err := scanner.Scan(urls[i], 2)
		if err != nil {
			return nil, err
		}

		for _, d := range data {
			res = append(res, d)
		}

		fmt.Printf(" done!\n")
	}

	return res, nil
}

func printResults(res map[string]string) {
	fmt.Printf("\n%v results\n", len(res))
	for k, v := range res {
		fmt.Println()
		fmt.Println(k)
		fmt.Println(v)
	}
}
