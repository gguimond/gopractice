package main

import (
    "fmt"
)

func sortedSquares(A []int) []int {
    var ipos, n int = 0, len(A)

    for ; A[ipos] < 0 && ipos < n; ipos++ {}

    ineg := ipos - 1
    result := make([]int, len(A))
    ir := 0

    for ; ipos < n && ineg >= 0; {
        if(A[ipos] * A[ipos] < A[ineg] * A[ineg]){
            result[ir] = A[ipos] * A[ipos]
            ir++
            ipos++
        }else{
            result[ir] = A[ineg] * A[ineg]
            ir++
            ineg--
        }
        
    }

    for ; ipos < n; {
        result[ir] = A[ipos] * A[ipos]
        ir++
        ipos++
    }

    for ; ineg >= 0; {
        result[ir] = A[ineg] * A[ineg]
        ir++
        ineg--
    }

    return result
}

func main() {
    a := []int{-4,-1,0,3,10}
    fmt.Printf("%v", sortedSquares(a));
}