package main

import (
    "fmt"
    "encoding/json"
)

var DefaultValue int = -1024

type TreeNode struct {
    Value int
    Left *TreeNode
    Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
    if(root == nil){
        return root
    }
    var left *TreeNode = invertTree(root.Left)
    var right *TreeNode = invertTree(root.Right)
    root.Left = right
    root.Right = left
    return root
}

func invertTreeIter(root *TreeNode) *TreeNode {
    if(root == nil){
        return root
    }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)

    for ; len(queue) > 0; {
        t := queue[0]
        queue = queue[1:]
        tmp := t.Right
        t.Right = t.Left
        t.Left = tmp
        if(t.Left != nil) {
            queue = append(queue, t.Left)
        }
        if(t.Right != nil) {
            queue = append(queue, t.Right)
        }
    }
    return root
}

func main() {
    treeNode1 := &TreeNode{
        Value: 4,
        Left: &TreeNode{
            Value: 2,
            Left: &TreeNode{
                Value: 1,
            },
            Right: &TreeNode{
                Value: 3,
            },
        },
        Right: &TreeNode{
            Value: 7,
            Left: &TreeNode{
                Value: 6,
            },
            Right: &TreeNode{
                Value: 9,
            },
        },
    }
    data, _ := json.Marshal(treeNode1)
    fmt.Printf("%s\n", data)

    invertedT := invertTree(treeNode1)
    data, _ = json.Marshal(invertedT)
    fmt.Printf("%s\n", data)

    invertedT = invertTree(treeNode1)

    invertedT2 := invertTreeIter(treeNode1)
    data, _ = json.Marshal(invertedT2)
    fmt.Printf("%s\n", data)
}