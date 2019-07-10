package main

import (
    "fmt"
)

func reverseBits(num uint32) uint32 {
    var shift, result uint32 = 31, 0

    for ; shift > 0 && num != 0; {
        if(num & 1 == 1){
            result |= (1 << shift)
        }
        num >>= 1
        shift --
    }

    return result

}

func main() {
    fmt.Printf("%v", reverseBits(43261596));
}