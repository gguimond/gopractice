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

func clumsyStack(N int) int {
    var stack []int
    var res int = 0
    var op int = 0

    stack = append(stack, N)
    for i:=N-1; i > 0; i, op = i-1, op+1 %4 {
        switch(op) {
        case 0:
            n := len(stack)-1
            el := stack[n]
            stack = stack[:n]
            stack = append(stack, el * i)
        case 1:
            n := len(stack)-1
            res += stack[n]/i
            stack = stack[:n]
        case 2:
            res += i
        case 3:
            stack = append(stack, -i)
        default:
        }
    }
    
    l := len(stack)
    for i := l-1; i >= 0; i-- {
        res += stack[i]
    }
    return res
}

func main() {
    fmt.Printf("%v", clumsy(4))
    fmt.Printf("%v", clumsy(10))
    fmt.Printf("%v", clumsyStack(4))
    fmt.Printf("%v", clumsyStack(10))
}