package main

import (
    "fmt"
)

func max(a, b int) int {
    if b <= a {
        return a
    }
    return b
}

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
    n := len(grid)
    maxRow, maxCol :=  make([]int, n), make([]int, n)
    var result int = 0
    for x:=0; x < n; x++ {
        for y:=0; y < n; y++ {
            maxRow[x] = max(maxRow[x], grid[x][y])
            maxCol[y] = max(maxCol[y], grid[x][y])
        }
    }
    fmt.Printf("%v", maxRow);
    fmt.Printf("%v", maxCol);

    for x:=0; x < n; x++ {
        for y:=0; y < n; y++ {
            result += min(maxRow[x], maxCol[y]) - grid[x][y]
        }
    }

    return result
}

func main() {
    grid := [][]int{{3,0,8,4}, {2,4,5,7}, {9,2,6,3}, {0,3,1,0}}
    fmt.Printf("%v", maxIncreaseKeepingSkyline(grid));
}