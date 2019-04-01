package main

import (
    "fmt"
    "sync"
    "errors"
)

// fetched tracks URLs that have been (or are being) fetched.
// The lock must be held while reading from or writing to the map.
// See http://golang.org/ref/spec#Struct_types section on embedded types.
var fetched = struct {
    m map[string]error
    sync.Mutex
}{m: make(map[string]error)}

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

var loading = errors.New("url load in progress") // sentinel value

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
    // TODO: Fetch URLs in parallel.
    // TODO: Don't fetch the same URL twice.
    // This implementation doesn't do either:
    if depth <= 0 {
        return
    }

    fetched.Lock()
    if _, ok := fetched.m[url]; ok {
        fetched.Unlock()
        fmt.Printf("<- Done with %v, already fetched.\n", url)
        return
    }

    // We mark the url to be loading to avoid others reloading it at the same time.
    fetched.m[url] = loading
    fetched.Unlock()

    body, urls, err := fetcher.Fetch(url)

    // And update the status in a synced zone.
    fetched.Lock()
    fetched.m[url] = err
    fetched.Unlock()
    
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("found: %s %q\n", url, body)
    done := make(chan bool)
    for _, u := range urls {
        go func(url string){
            Crawl(url, depth-1, fetcher)
            done <- true
        }(u)
    }
    for i := range urls{
        fmt.Printf("waiting for %v\n", i)
        <- done
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
