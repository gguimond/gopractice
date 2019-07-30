package main

import (
    "fmt"
    "sync"
)

func foo(n int, ch chan int) {
    for i:= 0; i < n; i++ {
        fmt.Println("foo")
        ch <- 1
        <- ch
    }    
}

func bar(n int, ch chan int) {
    for i:= 0; i < n; i++ {
        <- ch
        fmt.Println("bar")
        ch <- 1
    }    
}

func main() {
    var wg sync.WaitGroup
    var n int = 2
    wg.Add(2)
    ch := make(chan int)
    go func() {
        foo(n, ch)
        wg.Done()
    }()
    go func() {
        bar(n, ch)
        wg.Done()
    }()
    wg.Wait()
    close(ch)
}