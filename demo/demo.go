package main

import (
        "log"
        "net/http"
        "regexp"
        "sync"
        "bytes"
        "strconv"
)

var visitors struct {
    sync.Mutex
    n int
  }

//var colorRx = regexp.MustCompile(`^\w*$`)

var colorRxPool = sync.Pool{
        New: func() interface{} { return regexp.MustCompile(`\w*$`) },
}

var bufPool = sync.Pool{
        New: func() interface{} {
                return new(bytes.Buffer)
        },
}

func handleHi(w http.ResponseWriter, r *http.Request) {
        if !colorRxPool.Get().(*regexp.Regexp).MatchString(r.FormValue("color")){
                http.Error(w, "Optional color is invalid", http.StatusBadRequest)
                return
        }
        visitors.Lock()
        visitors.n++
        yourVisitorNumber := visitors.n
        visitors.Unlock()
        //yourVisitorNumber := nextVisitorNum()
        //w.Header().Set("Content-Type", "text/html; charset=utf-8")
        buf := bufPool.Get().(*bytes.Buffer)
        defer bufPool.Put(buf)
        buf.Reset()
        buf.WriteString("<h1 style='color: ")
        buf.WriteString(r.FormValue("color"))
        buf.WriteString(">Welcome!</h1>You are visitor number ")
        b := strconv.AppendInt(buf.Bytes(), int64(yourVisitorNumber), 10)
        b = append(b, '!')
        w.Write(b)
        //fmt.Fprintf(w, "<h1 style='color: %s'>Welcome!</h1>You are visitor number %v !", r.FormValue("color"), yourVisitorNumber)
}

func nextVisitorNum() int {
    visitors.Lock()
    defer visitors.Unlock()
    visitors.n++
    return visitors.n
}

func main() {
        log.Printf("Starting on port 8080")
        http.HandleFunc("/hi", handleHi)
        log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}