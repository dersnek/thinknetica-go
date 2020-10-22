package main

import (
	"fmt"
	"log"
	"strings"

	"dersnek/task-2/pkg/spider"
)

func main() {
	urls := []string{"https://katata.games", "https://go.dev"}
	const depth = 2
	store := make(map[string]string)
	var q string

	for i := 0; i < len(urls); i++ {
		fmt.Printf("Scanning %s...", urls[i])

		data, err := spider.Scan(urls[i], depth)
		if err != nil {
			log.Printf("error while scanning %s: %v\n", urls[i], err)
			return
		}

		for k, v := range data {
			store[k] = v
		}

		fmt.Printf(" done!\n")
	}

	for {
		q = ""

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
