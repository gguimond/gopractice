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

type Vertex struct {
    X int
    Y int
}

type Vertex2 struct {
    Lat, Long float64
}

func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
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

    i := 42
    p := &i         // point to i
    fmt.Println(*p) // read i through the pointer
    *p = 21         // set i through the pointer
    fmt.Println(i)  // see the new value of i

    v := Vertex{1, 2}
    v2 := Vertex{X: 1} 
    v.X = 4
    v2.X = 4
    fmt.Println(v.X)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes[0])
    var s []int = primes[1:4]
    fmt.Println(s)

    primes[1] = 1
    fmt.Println(s)

    r := []bool{true, false, true, true, false, true}
    fmt.Println(r)

    r = r[:3]
    fmt.Println(r)

    b := make([]int, 0, 5)
    fmt.Println(b)

    c := b[:2]
    fmt.Println(c)

    board := [][]string{
        []string{"_", "_", "_"},
        []string{"_", "_", "_"},
        []string{"_", "_", "_"},
    }

    board[0][0] = "X"

    var s2 []int
    s2 = append(s, 2, 3, 4)
    fmt.Println(s2)

    for v := range primes {
        fmt.Printf("%d\n", v)
    }

    var m map[string]Vertex2
    m = make(map[string]Vertex2)
    m["Bell Labs"] = Vertex2{
        40.68433, -74.39967,
    }
    fmt.Println(m)

    var m2 = map[string]Vertex2{
        "Bell Labs": Vertex2{
            40.68433, -74.39967,
        },
        "Google": Vertex2{
            37.42202, -122.08408,
        },
    }
    fmt.Println(m2)

    m3 := make(map[string]int)

    m3["foo"] = 1
    delete(m3, "foo")
    val, ok := m3["foo"]
    fmt.Println("The value:", val, "Present?", ok)

    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    fmt.Println(compute(hypot))

    adderx := adder()
    adderx2 := adder()
    fmt.Println(adderx(3))
    fmt.Println(adderx2(15))
}