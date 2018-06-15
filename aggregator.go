package scrapping

import (
	"sync"
)

type Aggregator struct {
	SearchEngines []SearchEngine
}

func (aggregator *Aggregator) Query(query string) []Result {
	merged := aggregator.merge(query)
	// all results -> merged
	return group(merged)
}

func (aggregator *Aggregator) merge(query string) <-chan Result {
	var combined = make(chan Result)
	var wg sync.WaitGroup
	wg.Add(len(aggregator.SearchEngines))
	for _, engine := range aggregator.SearchEngines {
		c := engine.Query(query)
		go func() {
			for result := range c {
				combined <- result
			}
			wg.Done()
		}()
	}
	// close channel
	go func() {
		wg.Wait()
		close(combined)
	}()
	return combined
}

// combine by url, keeping engine in sources field
func group(combined <-chan Result) []Result {

	m := make(map[string]Result)
	for next := range combined {
		existing, ok := m[next.url]
		if ok {
			next.sources = append(existing.sources, next.sources[0])
		}
		m[next.url] = next

	}

	values := make([]Result, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}

	return values
}
