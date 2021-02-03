package main

import (
	"bufio"
	"dersnek/task-6/pkg/engine"
	"fmt"
	"os"
)

func main() {
	engine := engine.New()
	err := engine.Start()
	if err != nil {
		fmt.Println(err)
		return
	}

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
		res := engine.Find(q)
		printResults(res)
	}
}

func printResults(res map[string]string) {
	fmt.Printf("\n%v results\n", len(res))
	for k, v := range res {
		fmt.Printf("\n%v\n%v\n", k, v)
	}
}
