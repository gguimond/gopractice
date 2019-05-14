package main

import (
    "context"
    "time"
    "fmt"
    "sync"
)

func operation(ctx context.Context) {
    fmt.Println("call")
    wg := ctx.Value("wg").(*sync.WaitGroup)
    defer wg.Done()
    select {
        case <-time.After(500 * time.Millisecond):
            fmt.Println("done")
        case <-ctx.Done():
            fmt.Println("halted operation")
    }
}

func main() {
    var wg sync.WaitGroup
    ctx := context.Background()
    wg.Add(3)

    ctx1 := context.WithValue(ctx, "wg", &wg)
    ctx2 := context.WithValue(ctx, "wg", &wg)
    ctx3 := context.WithValue(ctx, "wg", &wg)

    ctx2, cancel := context.WithCancel(ctx2)

    ctx3, _ = context.WithTimeout(ctx3, 30*time.Millisecond)

    go func(){
        operation(ctx1)
    }()

    go func(){
        operation(ctx2)
    }()

    go func(){
        operation(ctx3)
    }()

    cancel()
    wg.Wait()
}