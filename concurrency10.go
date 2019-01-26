//https://tour.golang.org/concurrency/10
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

type SafeCache struct {
    urls   map[string]int
    mux sync.Mutex
}

var visited = SafeCache{urls: make(map[string]int)}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
    if depth <= 0 {
        return
    }
    
    _, ok := visited.urls[url]
    if !ok {
        visited.mux.Lock()
        visited.urls[url] = 1
        visited.mux.Unlock()
        ch1 := make(chan string)
        ch2 := make(chan []string)
        ch3 := make(chan error)    
        go func(ch1 chan string, ch2 chan []string, ch3 chan error){
            body, urls, err := fetcher.Fetch(url)
            ch1<-body
            ch2<-urls
            ch3<-err
        }(ch1, ch2, ch3)
        body, urls, err := <-ch1, <-ch2, <-ch3
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("found: %s %q\n", url, body)
        for _, u := range urls {
            Crawl(u, depth-1, fetcher)
        }
    }
    return
}

func main() {
    Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
    if res, ok := f[url]; ok {
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
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