package main

import (
    "fmt"
    "sync"
)

func foo(n int, m1 *sync.Mutex, m2 *sync.Mutex) {
    for i:= 0; i < n; i++ {
        m1.Lock()
        fmt.Println("foo")
        m2.Unlock()
    }    
}

func bar(n int, m1 *sync.Mutex, m2 *sync.Mutex) {
    for i:= 0; i < n; i++ {
        m2.Lock()
        fmt.Println("bar")
        m1.Unlock()
    }    
}

func main() {
    var wg sync.WaitGroup
    var n int = 2
    wg.Add(2)
    var m1, m2 sync.Mutex
    m2.Lock()
    go func() {
        defer wg.Done()
        foo(n, &m1, &m2)
    }()
    go func() {
        defer wg.Done()
        bar(n, &m1, &m2)
    }()
    wg.Wait()
}