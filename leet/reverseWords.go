package main

import (
    "fmt"
    "strings"
)

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func reverseWords(s string) string {
    words := strings.Split(s, " ")
    var result strings.Builder

    for i := 0; i < len(words); i++ {
        result.WriteString(reverse(words[i]))
        result.WriteString(" ")
    }

    return strings.Trim(result.String(), " ")
}

func main() {
    fmt.Printf("%s", reverseWords("Let's take LeetCode contest"));
    fmt.Printf("%s", "end")
}