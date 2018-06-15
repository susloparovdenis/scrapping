package scrapping

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

type SearchEngine interface {
	Query(query string) <-chan Result
}

type WebEngine struct {
	engineUrl, entrySelector, titleSelector, name string
	getUrl                                        func(*goquery.Selection) string
}

func (engine *WebEngine) Query(query string) <-chan Result {
	// Request the HTML page.
	results := make(chan Result)
	go func() {
		queryUrl := fmt.Sprintf(engine.engineUrl, url.QueryEscape(query))
		res, err := http.Get(queryUrl)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		find := doc.Find(engine.entrySelector)
		find.Each(func(i int, s *goquery.Selection) {
			href := engine.getUrl(s)
			results <- Result{
				s.Find(engine.titleSelector).Text(),
				href,
				[]string{engine.name}}

		})
		close(results)
	}()
	return results
}

var googleUrlRegex = regexp.MustCompile(`q=(.*?)&`)

func NewGoogle() *WebEngine {
	getUrl := func(selection *goquery.Selection) (url string) {
		dirty, _ := selection.Find(".r>a").Attr("href")
		url = googleUrlRegex.FindStringSubmatch(dirty)[1]
		return

	}
	return &WebEngine{
		engineUrl:     "https://www.google.nl/search?q=%s",
		entrySelector: ".g",
		titleSelector: ".r>a",
		name:          "google",
		getUrl:        getUrl}
}

var yahooUrlRegex = regexp.MustCompile(`=(https|http.*)/RK`)

func NewYahoo() *WebEngine {
	getUrl := func(selection *goquery.Selection) (href string) {
		href, _ = selection.Find("a.ac-algo").Attr("href")
		matches := yahooUrlRegex.FindStringSubmatch(href)
		if len(matches) == 2 {
			href, _ = url.QueryUnescape(matches[1])
		}
		return
	}
	return &WebEngine{
		engineUrl:     "https://search.yahoo.com/search?p=%s",
		entrySelector: ".algo",
		titleSelector: ".title > a",
		name:          "yahoo",
		getUrl:        getUrl}
}
