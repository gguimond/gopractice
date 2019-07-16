package main

import (
    "fmt"
    "encoding/json"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    if(head == nil || head.Next == nil) {
        return head 
    }
    p := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil

    return p
}

func reverseListIter(head *ListNode) *ListNode {
    if(head == nil || head.Next == nil) {
        return head 
    }
    var cur, prev *ListNode 
    cur = head
    for ; cur != nil; {
        next := cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }

    return prev
}

func main() {
    l := &ListNode{
        Val: 1,
        Next: &ListNode{
            Val: 2,
            Next: &ListNode{
                Val: 3,
            },
        },
    }

    l2 := &ListNode{
        Val: 1,
        Next: &ListNode{
            Val: 2,
            Next: &ListNode{
                Val: 3,
            },
        },
    }

    data, _ := json.Marshal(reverseList(l))
    fmt.Printf("%s\n", data)

    data, _ = json.Marshal(reverseList(l2))
    fmt.Printf("%s\n", data)
}