package main

import (
    "fmt"
)

func isPerfectSquare(num int) bool {

    if(num == 0 || num ==1){
        return true
    }

    for left, right := 1, num; left <= right; {
        mid := (right - left)/2 + left
        if(mid*mid == num){
            return true
        }

        if(mid*mid > num){
            right = mid - 1
        }else{
            left = mid + 1
        }
    }
    return false
}

func main() {
    fmt.Printf("%v", isPerfectSquare(14));
    fmt.Printf("%v", isPerfectSquare(16));
}