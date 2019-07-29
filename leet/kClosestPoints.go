package main

import (
    "fmt"
    "math/rand"
)

func kClosest(points [][]int, K int) [][]int {
    sort(&points, 0, len(points) - 1, K)
    return points[0:K]
}

func sort(points *[][]int, i int, j int, K int) {
    if j <= i {
        return
    }
    k := rand.Intn(j - i) + i
    swap(points, i, k)

    mid := partition(points, i, j)
    leftSize := mid - i + 1
    if K < leftSize {
        sort(points, 0, mid - 1, K)
    } else {
        sort(points, mid + 1, j, K - leftSize)
    }
}

func partition(points *[][]int, i int, j int) int {
    oi := i
    pivotDist := dist((*points)[i])
    i++
    for {
        for dist((*points)[i]) < pivotDist && i < j{
            i++
        }
        for dist((*points)[j]) > pivotDist && i <= j{
            j++
        }
        if (i >= j) {
            break
        }
        swap(points, i,j)
    }
    swap(points, oi,j)
    return j
}

func swap(points *[][]int, i int, j int) {
    var tmp []int
    tmp = (*points)[i]
    (*points)[i] = (*points)[j]
    (*points)[j] = tmp
}

func dist(a []int) int {
    return a[0] * a[0] + a[1] * a[1]
}

func main() {
    fmt.Printf("%v", kClosest([][]int{{1,3},{-2,2}}, 1))
    fmt.Println("")
    fmt.Printf("%v", kClosest([][]int{{3,3},{5,-1},{-2,4}}, 2))
}