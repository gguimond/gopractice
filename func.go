package main

import (
    "fmt"
)

type checker func(s string) (isOk bool)

type checkerProvider func() checker

func letPassEverything(s string) (isOk bool){
    return true;
}

func discardEverything(s string) (isOk bool){
    return false;
}

func myfunc(k int) checkerProvider {
    return func() checker {
        if k > 0 {
            return discardEverything
        }
        return letPassEverything
    }
}

func main() {
    pass := myfunc(0)
    discard := myfunc(1)
    fmt.Printf("pass : %v\n", pass()("test"));
    fmt.Printf("discard : %v\n", discard()("test"));
}