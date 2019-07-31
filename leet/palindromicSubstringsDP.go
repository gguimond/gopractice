package main

import (
    "fmt"
)

func countSubstrings(s string) int {
    n := len(s)
    result := 0
    table := make([][]bool, n)
    for i := range table { 
        table[i] = make([]bool, n)
        table[i][i] = true
        result++
        if i < n-1 && s[i] == s[i+1] {
            table[i][i+1] = true
            result++
        }
    }
    var j int
    for k:=2; k < n; k++ {
        for i:=0; i < n-k; i++ {
            j = i+k
            if table[i+1][j-1] && s[i] == s[j] {
                table[i][j] = true
                result++
            }
        }
    }
    fmt.Printf("%v", table)
    return result
}

func main() {
    fmt.Printf("%v", countSubstrings("kayak"))
    fmt.Printf("%v", countSubstrings("abc"))
    fmt.Printf("%v", countSubstrings("aaa"))
}