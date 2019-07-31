package main

import (
    "fmt"
)

func countSubstrings(s string) int {
    n := len(s)
    result := 0
    var left, right int
    for i:=0; i < 2*n - 1; i++ {
        left = i/2
        right = left + i % 2
        for left >= 0 && right < n && s[left] == s[right] {
            result++
            right++
            left--
        }
    }
    return result
}

func main() {
    fmt.Printf("%v", countSubstrings("kayak"))
    fmt.Printf("%v", countSubstrings("abc"))
    fmt.Printf("%v", countSubstrings("aaa"))
}