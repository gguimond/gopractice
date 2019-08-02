package main

import (
    "fmt"
    "strings"
)

//Set

var exists = struct{}{}

type set struct {
    m map[string]struct{}
}

func NewSet() *set {
    s := &set{}
    s.m = make(map[string]struct{})
    return s
}

func (s *set) Add(value string) {
    s.m[value] = exists
}

func (s *set) Remove(value string) {
    delete(s.m, value)
}

func (s *set) Contains(value string) bool {
    _, c := s.m[value]
    return c
}

//Trie

type (
    Trie struct {
        root *node
        size int
    }
    node struct {
        key interface{}
        value interface{}
        next [256]*node
    }
    iterator struct {
        step int
        node *node
        prev *iterator
    }
)

func toBytes(obj interface{}) []byte {
    switch o := obj.(type) {
    case []byte:
        return o
    case string:
        return []byte(o)
    }
    return []byte(fmt.Sprint(obj))
}

func NewTrie() *Trie {
    return &Trie{nil,0}
}

func newNode() *node {
    var next [256]*node
    return &node{nil,nil,next}
}

func (this *Trie) Insert(key interface{}, value interface{}) {
    if this.size == 0 {
        this.root = newNode()
    }
    
    bs := toBytes(key)
    cur := this.root
    for i := 0; i < len(bs); i++ {
        if cur.next[bs[i]] != nil {
            cur = cur.next[bs[i]]
        } else {
            cur.next[bs[i]] = newNode()
            cur = cur.next[bs[i]]
        }
    }   
    if cur.key == nil {
        this.size++
    }
    cur.key = key
    cur.value = value
}


func replaceWords(dict []string, sentence string) string {
    var set = NewSet()
    var result strings.Builder
    for _, root := range dict {
        set.Add(root)
    }

    for _, word := range strings.Split(sentence, " ") {
        w := ""
        for i:= 0; i < len(word); i++ {
            w = word[:i+1]
            if(set.Contains(word[:i+1])) {
                break
            }
        }
        result.WriteString(w)
        result.WriteString(" ")
    }
    return strings.Trim(result.String(), " ")
}

func replaceWordsTrie(dict []string, sentence string) string {
    var tr = NewTrie()
    var result strings.Builder
    for _, root := range dict {
        tr.Insert(root, struct{}{})
    }

    for _, word := range strings.Split(sentence, " ") {
        cur := tr.root
        for _, c := range word {
            if(cur.next[c] == nil) {
                break;
            } else {
                cur = cur.next[c]
            }
        }
        var w string
        if(cur.key != nil) {
            w = cur.key.(string)
        } else {
            w = word
        }
        result.WriteString(w)
        result.WriteString(" ")
    }
    return strings.Trim(result.String(), " ")
}

func main() {
    fmt.Printf("%v", replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
    fmt.Printf("%v", replaceWordsTrie([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
}