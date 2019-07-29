package main

import (
    "fmt"
    "regexp"
    "strconv"
)

func complexNumberMultiply(a string, b string) string { 
    as := regexp.MustCompile("(?P<real>[^+]*)\\+(?P<complex>[^i]*)i").FindStringSubmatch(a)
    bs := regexp.MustCompile("(?P<real>[^+]*)\\+(?P<complex>[^i]*)i").FindStringSubmatch(b)
    aReal, _ := strconv.ParseInt(as[1], 10, 8)
    aComplex, _ := strconv.ParseInt(as[2], 10, 8)
    bReal, _ := strconv.ParseInt(bs[1], 10, 8)
    bComplex, _ := strconv.ParseInt(bs[2], 10, 8)
    
    return strconv.FormatInt(aReal*bReal - aComplex*bComplex, 10) + "+" + strconv.FormatInt(aComplex*bReal + aReal*bComplex, 10) + "i"
}

func main() {
    fmt.Printf("%v", complexNumberMultiply("1+1i", "1+1i"));
    fmt.Println("")
    fmt.Printf("%v", complexNumberMultiply("1+-1i", "1+-1i"));
}