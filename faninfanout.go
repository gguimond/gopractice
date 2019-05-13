package main

import (
    "fmt"
    "sync"
)

func main() {
    done := make(chan struct{})
    defer close(done)

    // Set up the pipeline.
    c := gen(done, 2, 3)
    c1 := sq(done, c)
    c2 := sq(done, c)

    
    for n := range merge(done, c1, c2) {
        fmt.Println(n)
    }

}

func gen(done <-chan struct{}, nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            select {
                case out <- n:
                case <- done:
                    return
            }
        }
    }()
    return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            select {
                case out <- n * n:
                case <- done:
                    return
            }
        }
    }()
    return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int{
    var wg sync.WaitGroup
    out := make(chan int)

    output := func(c <-chan int) {
        defer wg.Done()
        for n := range c {
            select {
                case out <- n:
                case <-done:
                    return
            }
        }
    }

    wg.Add(len(cs))

    for _, c := range cs {
        go output(c)
    }

    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}