package main

import (
    "fmt"
)

func clumsy(N int) int {
    switch(N) {
    case 0:
        return 0
    case 1:
        return 1
    case 2:
        return 2
    case 3:
        return 6
    case 4:
        return 7
    default: 
        return doClumsy(N)
    }
}

func doClumsy(N int) int {
    res := N * (N-1) / (N-2) + (N-3)
    next := (N-4) * max(1, (N-5)) / max(1, (N-6))  
    return res - 2*next + clumsy(N-4)
}

func max(a, b int) int {
    if b <= a {
        return a
    }
    return b
}

func main() {
    fmt.Printf("%v", clumsy(4))
    fmt.Printf("%v", clumsy(10))
}