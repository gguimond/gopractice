package main

import (
    "testing"
)

var result int

func BenchmarkClumsy10(b *testing.B) {
    var r int
    // run the Fib function b.N times
    for n := 0; n < b.N; n++ {
        r = clumsy(10)
    }
    result = r
}

func BenchmarkClumsyStack10(b *testing.B) {
    var r int
    // run the Fib function b.N times
    for n := 0; n < b.N; n++ {
        r = clumsyStack(10)
    }
    result = r
}

func BenchmarkClumsy100(b *testing.B) {
    var r int
    // run the Fib function b.N times
    for n := 0; n < b.N; n++ {
        r = clumsy(100)
    }
    result = r
}

func BenchmarkClumsyStack100(b *testing.B) {
    var r int
    // run the Fib function b.N times
    for n := 0; n < b.N; n++ {
        r = clumsyStack(100)
    }
    result = r
}