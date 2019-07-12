package main

import (
    "fmt"
    "strings"
)

func convert(s string, numRows int) string {
    var currentRow int = 0
    var goDown bool = false
    var nbRows int
    var result strings.Builder

    if(numRows == 1) {
        return s
    }

    if(numRows > len(s)) {
        nbRows = len(s) 
    } else {
        nbRows = numRows
    }

    res := make([][]byte, nbRows)

    for i := 0; i < len(s); i++ {
        res[currentRow] = append(res[currentRow], s[i])
        if(currentRow == 0 || currentRow == nbRows -1){
            goDown = !goDown
        }
        if(goDown){
            currentRow++
        } else{
            currentRow--
        }
    }

    for i:=0; i < len(res); i++ {
        result.WriteString(string (res[i]))
    }

    return result.String()
}

func main() {
    fmt.Printf("%s", convert("PAYPALISHIRING", 3));
    fmt.Println("\n")
    fmt.Printf("%s", convert("PAYPALISHIRING", 4));
}