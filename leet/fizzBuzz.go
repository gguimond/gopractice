package main

import (
    "fmt"
    "strings"
    "strconv"
)

func fizzBuzz(n int) []string {
    result := []string{}
    var str strings.Builder
    for i := 1; i <= n; i++ {
        str.Reset()
        if(i%3 == 0){
            str.WriteString( "Fizz")
        }
        if(i%5 == 0){
            str.WriteString( "Buzz")
        }
        if(i%3 != 0 && i%5 != 0){
            result = append(result, strconv.Itoa(i))
        }else{
            result = append(result, str.String())
        }
    }
    return result
}

func main() {
    fmt.Printf("%s", fizzBuzz(15));
}