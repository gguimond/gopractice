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

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if(t1 == nil){
        return t2
    }
    if(t2 == nil){
        return t1
    }
    t1.Value += t2.Value
    t1.Left = mergeTrees(t1.Left, t2.Left)
    t1.Right = mergeTrees(t1.Right, t2.Right)
    return t1
}

func main() {
    treeNode1 := &TreeNode{
        Value: 1,
        Left: &TreeNode{
            Value: 2,
            Left: &TreeNode{
                Value: 4,
            },
            Right: &TreeNode{
                Value: 5,
            },
        },
        Right: &TreeNode{
            Value: 3,
        },
    }
    data, _ := json.Marshal(treeNode1)
    fmt.Printf("%s\n", data)

    treeNode2 := &TreeNode{
        Value: 5,
        Left: &TreeNode{
            Value: 3,
            Left: &TreeNode{
                Value: 2,
            },
        },
        Right: &TreeNode{
            Value: 6,
            Left: &TreeNode{
                Value: 1,
            },
            Right: &TreeNode{
                Value: 4,
            },
        },
    }
    data, _ = json.Marshal(treeNode2)
    fmt.Printf("%s\n", data)

    mergedT := mergeTrees(treeNode1, treeNode2)
    data, _ = json.Marshal(mergedT)
    fmt.Printf("%s\n", data)
}