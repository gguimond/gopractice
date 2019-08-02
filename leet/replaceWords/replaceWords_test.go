package main

import (
    "testing"
)

var result string

func BenchmarkReplaceWords(b *testing.B) {
    var r string
    for n := 0; n < b.N; n++ {
        r = replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery")
    }
    result = r
}

func BenchmarkReplaceWordsTrie(b *testing.B) {
    var r string
    for n := 0; n < b.N; n++ {
        r = replaceWordsTrie([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery")
    }
    result = r
}

func BenchmarkReplaceWordsBigSentence(b *testing.B) {
    var r string
    for n := 0; n < b.N; n++ {
        r = replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery the cattle was rattled by the battery the cattle was rattled by the battery the cattle was rattled by the battery the cattle was rattled by the battery")
    }
    result = r
}

func BenchmarkReplaceWordsTrieBigSentence(b *testing.B) {
    var r string
    for n := 0; n < b.N; n++ {
        r = replaceWordsTrie([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery the cattle was rattled by the battery the cattle was rattled by the battery the cattle was rattled by the battery the cattle was rattled by the battery")
    }
    result = r
}

func BenchmarkReplaceWordsBigDict(b *testing.B) {
    var r string
    for n := 0; n < b.N; n++ {
        r = replaceWords([]string{"cat", "bat", "rat", "aa", "aaa", "aaaa"}, "the cattle was rattled by the battery")
    }
    result = r
}

func BenchmarkReplaceWordsTrieBigDict(b *testing.B) {
    var r string
    for n := 0; n < b.N; n++ {
        r = replaceWordsTrie([]string{"cat", "bat", "rat", "aa", "aaa", "aaaa"}, "the cattle was rattled by the battery")
    }
    result = r
}