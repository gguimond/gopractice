package main

import (
    "fmt"
)

func sortArrayByParityII(A []int) []int {
    var i, j int = 0, 1

    for ; i < len(A) ; i += 2 {
        if(A[i]%2 == 1){
            for ;A[j]%2 == 1; j += 2{
            }
            tmp := A[i]
            A[i] = A[j]
            A[j] = tmp
        }
    }
    return A
}

func main() {
    a := []int{1,0,2,11,17,16}
    fmt.Printf("%v", sortArrayByParityII(a));
}