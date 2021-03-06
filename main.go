package main

import (
	"fmt"
	"io"
	"math"
	"runtime"
	"strings"
	"sync"
	"time"
    "strconv"
    "log"
    "os"

	"github.com/gguimond/gopractice/mylib"
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

func (v *Vertex) Abs() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

type Abser interface {
	Abs() float64
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

var (
	test string = "toto"
)

func returnBoolAndInt() (a, b int) {
	a = 1
	b = 2
	return
}

func main() {
	defer fmt.Println("defer")
	defer fmt.Println("will be called first")

	x := "short"
	fmt.Println("Hello, world.")
	fmt.Println(mylib.Reverse("toto"))
	fmt.Println(swap("toto", "titi"))
	fmt.Println(add(1, 2))
	fmt.Println(split(17))
	fmt.Println(add(1, 2))
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

	vobj := Vertex{1, 2}
	fmt.Println(vobj.Abs())

	var ab Abser
	vInter := Vertex{3, 4}
	ab = &vInter
	fmt.Println(ab.Abs())

	vInter2, ok := ab.(*Vertex)
	fmt.Println(ok)
	fmt.Println(vInter2)

	switch ab.(type) {
	case *Vertex:
		fmt.Println("Vertex")
	default:
		// no match; here v has the same type as i
	}

	person := Person{"foobar", 18}
	fmt.Println(person)

	if err := run(); err != nil {
		fmt.Println(err)
	}

	rstr := strings.NewReader("Hello, Reader!")
	buffer := make([]byte, 8)
	for {
		nbytes, err := rstr.Read(buffer)
		fmt.Printf("n = %v err = %v b = %v\n", nbytes, err, buffer)
		fmt.Printf("b[:nbytes] = %q\n", buffer[:nbytes])
		if err == io.EOF {
			break
		}
	}

	ccount := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go ccount.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(ccount.Value("somekey"))

	var x1, y1 int = returnBoolAndInt()
	fmt.Printf("func return : %d, %d", x1, y1)

	var present bool
	_, present = timeZone["UTC"]
	fmt.Printf("found timezone : %d", present)
	_, present = timeZone["foo"]
	fmt.Printf("found timezone : %d", present)
	delete(timeZone, "UTC")

	mini := Min(1, 2, 3)
    mys := []string{"min : ", strconv.Itoa(mini)}
	fmt.Println(strings.Join(mys, " "))

    type ByteSize float64

    const (
        _           = iota // ignore first value by assigning to blank identifier
        KB ByteSize = 1 << (10 * iota)
        MB
        GB
        TB
        PB
        EB
        ZB
        YB
    )

    fmt.Println(KB)
    fmt.Println(MB)
    fmt.Println(1<<1)

    _ , ok17 := ab.(Abser)
    if ok17 {
        fmt.Println("is Abser")
    } else {
        fmt.Println("is not Abser")
    }

    job := Job{"print", log.New(os.Stderr, "Job: ", log.Ldate)}
    job.Println("starting now...")

    defer func() {
        if e := recover(); e != nil {
            _ = e.(OurError) // Will re-panic if not our error.
            fmt.Println("recover")
        }
    }()
    errorPanic()

}

type Job struct {
    Command string
    *log.Logger
}

func (job *Job) Printf(format string, args ...interface{}) {
    job.Logger.Printf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}

type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// ReadWriter is the interface that combines the Reader and Writer interfaces.
type ReadWriter interface {
    Reader
    Writer
}

func init(){
    fmt.Println("init")
}

func Min(a ...int) int {
	min := int(^uint(0) >> 1) // largest int
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

var timeZone = map[string]int{
	"UTC": 0 * 60 * 60,
	"EST": -5 * 60 * 60,
	"CST": -6 * 60 * 60,
	"MST": -7 * 60 * 60,
	"PST": -8 * 60 * 60,
}

func offset(tz string) int {
	if seconds, ok := timeZone[tz]; ok {
		return seconds
	}
	fmt.Println("unknown time zone:", tz)
	return 0
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type OurError string
func (e OurError) Error() string {
    return string(e)
}

func errorPanic() {
    panic(OurError("panic"))
}
