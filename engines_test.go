package scrapping

import (
	"github.com/kr/pretty"
	"testing"
)

func TestGoogle(t *testing.T) {
	google := NewGoogle()
	results := resultToSlice(google.Query("concurrency wiki"))
	t.Log(pretty.Sprint(results))
	if len(results) == 0 {
		t.Fatal("no results")
	}
}

func TestYahoo(t *testing.T) {
	yahoo := NewYahoo()
	results := resultToSlice(yahoo.Query("concurrency wiki"))
	t.Log(pretty.Sprint(results))
	if len(results) == 0 {
		t.Fatal("no results")
	}
}
