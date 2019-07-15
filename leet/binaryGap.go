package main

import (
    "fmt"
)

func binaryGap(N int) int {
    var result, last int = 0, -1
    N32 := uint32(N)

    for i:= 0; i < N; i++ {
        if(N32 >> uint32(i) & 1 == 1) {
            if(last >= 0 && i - last > result){
                result = i - last
            }
            last = i
        }
    } 
    return result
}

func main() {
    fmt.Printf("%v", binaryGap(22));
    fmt.Println("\n")
    fmt.Printf("%v", binaryGap(6));
    fmt.Println("\n")
    fmt.Printf("%v", binaryGap(8));
}