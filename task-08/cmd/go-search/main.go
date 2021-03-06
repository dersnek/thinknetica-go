package main

import (
	"bufio"
	"dersnek/task-8/pkg/crawler"
	"dersnek/task-8/pkg/crawler/spider"
	"dersnek/task-8/pkg/engine"
	"dersnek/task-8/pkg/index"
	"dersnek/task-8/pkg/index/hash"
	"dersnek/task-8/pkg/store"
	"dersnek/task-8/pkg/store/file"
	"fmt"
	"log"
	"os"
)

type gosearch struct {
	engine  *engine.Service
	scanner crawler.Interface
	index   index.Interface
	store   store.Interface

	sites []string
	depth int
}

func main() {
	server := new()
	server.init()
	server.run()
}

func new() *gosearch {
	gs := gosearch{
		sites: []string{"https://katata.games", "https://go.dev"},
		depth: 2,
	}
	gs.scanner = spider.New()
	gs.index = hash.New()
	gs.store = file.New()
	gs.engine = engine.New(gs.index)
	return &gs
}

func (gs *gosearch) init() {
	err := gs.engine.Start(gs.store, gs.scanner, gs.sites)
	if err != nil {
		log.Fatal("Failed to start the search engine:", err)
	}
}

func (gs *gosearch) run() {
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
		res := gs.engine.Find(q, gs.store)
		printResults(res)
	}
}

func printResults(res []store.Doc) {
	fmt.Printf("\n%v results\n", len(res))
	for _, p := range res {
		fmt.Printf("\n%v\n%v\n", p.URL, p.Title)
	}
}
