package scrapping

import (
	"github.com/kr/pretty"
	"testing"
)

type fakeEngine struct {
	results []Result
}

func (engine fakeEngine) Query(query string) <-chan Result {
	c := make(chan Result)
	go func() {
		for _, e := range engine.results {
			c <- e
		}
		close(c)
	}()
	return c
}

func TestMerge(t *testing.T) {
	e1 := fakeEngine{[]Result{
		{"u1", "u1", []string{"c1"}},
		{"u2", "u2", []string{"c1"}},
	}}
	e2 := fakeEngine{[]Result{
		{"u2", "u2", []string{"c2"}},
		{"u3", "u3", []string{"c2"}}},
	}
	aggregator := Aggregator{[]SearchEngine{e1, e2}}

	res := aggregator.merge("")
	for r := range res {
		t.Log(pretty.Sprint(r))
	}

}

func TestScrap(t *testing.T) {
	e1 := fakeEngine{[]Result{
		{"u1", "u1", []string{"c1"}},
		{"u2", "u2", []string{"c1"}},
	}}
	e2 := fakeEngine{[]Result{
		{"u2", "u2", []string{"c2"}},
		{"u3", "u3", []string{"c2"}}},
	}
	aggregator := Aggregator{[]SearchEngine{e1, e2}}

	res := aggregator.Query("")
	pretty.Log(res)

	if len(res) != 3 {
		t.Fatal("wrong number of results")
	}

}
