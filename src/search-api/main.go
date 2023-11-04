package main

import (
	"search-api/queue"
	"search-api/solr"
	"sync"
)

func main() {
	queue.InitQueue()
	solr.InitSolr()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		queue.Consume()
	}()

	// Wait for the goroutine to finish
	wg.Wait()
}
