package main

import (
    "fmt"
)

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
    result := [][]int{}
    queue := [][]int{}
    visited := map[int]map[int]bool{}

    queue = append(queue, []int{r0, c0})

    for ; len(queue) > 0; {
        p := queue[0]
        queue = queue[1:]

        r:= p[0]
        c:= p[1]

        if _, ok := visited[r][c]; ok {
            continue;
        }

        result = append(result, p)

        if r > 0 {
            queue = append(queue, []int{r - 1, c})
        }

        if c > 0 {
            queue = append(queue, []int{r, c - 1})
        }

        if r < R -1 {
            queue = append(queue, []int{r + 1, c})
        }

        if c < C -1 {
            queue = append(queue, []int{r, c + 1})
        }

        if _, ok := visited[r]; !ok {
            visited[r] = map[int]bool{}
        }
        visited[r][c] = true
    }

    return result
}

func main() {
    fmt.Printf("%v", allCellsDistOrder(2, 3, 1, 2));
}