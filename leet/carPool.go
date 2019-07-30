package main

import (
    "fmt"
)

func carPooling(trips [][]int, capacity int) bool {
    nbPassengersPerLocation := [1000]int{}

    for _, trip := range trips {
        nbPassengersPerLocation[trip[1]] += trip[0]
        nbPassengersPerLocation[trip[2]] -= trip[0]
    }

    currentPassenger := 0
    for _, l := range nbPassengersPerLocation {
        currentPassenger += l
        if currentPassenger > capacity {
            return false
        }
    }

    return true
}

func main() {
    fmt.Printf("%v", carPooling([][]int{{2,1,5}, {3,3,7}}, 4));
    fmt.Printf("%v", carPooling([][]int{{2,1,5}, {3,3,7}}, 5));
    fmt.Printf("%v", carPooling([][]int{{2,1,5}, {3,5,7}}, 3));
    fmt.Printf("%v", carPooling([][]int{{3,2,7}, {3,7,9}, {8,3,9}}, 11));
}