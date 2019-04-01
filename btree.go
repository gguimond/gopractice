package main

import (
    "golang.org/x/tour/tree"
    "fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    walkTree(t, ch)
    close(ch)
}

func walkTree(t *tree.Tree, ch chan int) {
    if t.Left != nil {
        walkTree(t.Left, ch)
    }
    ch <- t.Value
    if t.Right != nil {
        walkTree(t.Right, ch)
    }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    c1, c2 := make(chan int), make(chan int)
    go Walk(t1, c1)
    go Walk(t2, c2)
    for i := range c1 {
        if i != <- c2 {
            return false
        }
    }
    return true;
}

func main() {
    c := make(chan int)
    go Walk(tree.New(1), c)
    for i := range c {
        fmt.Println(i)
    }
    fmt.Println("Should return true:", Same(tree.New(1), tree.New(1)))
    fmt.Println("Should return false:", Same(tree.New(1), tree.New(2)))
}