package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}


func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // send sum to c
}


func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}

func fibonacci2(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

type Mystruct struct {
        c chan func(i int)
        i int
}

func (m *Mystruct) Multiply() {
    m.c <- func(i int) {
        m.i = m.i * i;
        fmt.Printf("multiply : %d", m.i)
        if(m.i >= 1024){
            fmt.Println("close")
            close(m.c)
        }
    }
}

func (m *Mystruct) loop () {
    /*for{
        select{
            case op := <- m.c:
                op(2)
            case <- m.quit:
                fmt.Println("quit")
                close(m.c)
                close(m.quit)
        }
    }*/
    for op := range m.c {
        op(2)
    }
}

func main() {
    go say("world")
    say("hello")

    s := []int{7, 2, 8, -9, 4, 0}
    c := make(chan int, 3)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c // receive from c

    fmt.Println(x, y, x+y)

    c2 := make(chan int, 10)
    go fibonacci(cap(c2), c2)
    for i := range c2 {
        fmt.Println(i)
    }

    c3 := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c3)
        }
        quit <- 0
    }()
    fibonacci2(c3, quit)

    /*tick := time.Tick(100 * time.Millisecond)
    boom := time.After(500 * time.Millisecond)
    for {
        select {
        case <-tick:
            fmt.Println("tick.")
        case <-boom:
            fmt.Println("BOOM!")
            return
        default:
            fmt.Println("    .")
            time.Sleep(50 * time.Millisecond)
        }
    }*/

    cFunc := make(chan func(i int))
    //cQuit := make(chan bool)
    mystruct := Mystruct{cFunc, 1}

    for z:=0; z < 10; z++{
        go mystruct.Multiply()
    }

    
    mystruct.loop()

    
    
}
