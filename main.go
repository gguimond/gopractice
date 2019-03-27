package main

import (
    "fmt"
    "github.com/gguimond/gopractice/mylib"
    "math"
    "runtime"
    "time"
)

func add(x, y int) int {
    return x + y
}

func swap(x, y string) (string, string) {
    return y, x
}

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

var foo string = "bar"

var a, b string = "yep", "yop"

const Pi = 3.14

const (
    // Create a huge number by shifting a 1 bit left 100 places.
    // In other words, the binary number that is 1 followed by 100 zeroes.
    Big = 1 << 100
    // Shift it right again 99 places, so we end up with 1<<1, or 2.
    Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
    return x * 0.1
}

func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    }
    return lim
}

func main() {
    defer fmt.Println("defer")
    defer fmt.Println("will be called first")

    x := "short"
    fmt.Println("Hello, world.")
    fmt.Println(mylib.Reverse("toto"))
    fmt.Println(swap("toto", "titi"))
    fmt.Println(add(1,2))
    fmt.Println(split(17))
    fmt.Println(add(1,2))
    fmt.Println(a)
    fmt.Println(x)
    fmt.Println(needInt(Small))
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))

    sum := 0    
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)

    fmt.Println(sqrt(2), sqrt(-4))

    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )

    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default:
        // freebsd, openbsd,
        // plan9, windows...
        fmt.Printf("%s.\n", os)
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")
    }
}