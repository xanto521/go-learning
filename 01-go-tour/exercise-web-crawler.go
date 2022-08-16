package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var urlsFlag = SafeUrlFlag{v: make(map[string]int)}

type SafeUrlFlag struct {
	v   map[string]int
	mux sync.Mutex
}

//get value
func (uf *SafeUrlFlag) Get(key string) int {
	uf.mux.Lock()
	defer uf.mux.Unlock()
	return uf.v[key]
}

//set value
func (uf *SafeUrlFlag) Set(key string, value int) {
	uf.mux.Lock()
	defer uf.mux.Unlock()
	uf.v[key] = value
}

func Crawl(url string, depth int, fetcher Fetcher) {
	//todo implement the crawl function here and then uncomment the following line
	//todo do not crawl the same URL twice
	if depth <= 0 {
		return
	}

	tmp := urlsFlag.Get(url)
	if tmp == 1 {
		fmt.Printf("this url is crawled: %s\n", url)
		return
	} else {
		urlsFlag.Set(url, 1)
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

type fakeFetcher map[string]*fakeResult
type fakeResult struct {
	body string
	urls []string
}

// 为什么这个fakeFetcher的fetch方法能加载到Fetcher接口中？
// 因为fakeFetcher实现了Fetcher接口，所以它的方法也能加载到Fetcher接口中。
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}
