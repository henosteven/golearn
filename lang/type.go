package main

import (
    "fmt"
)

type MyInt int

func main() {
    var j MyInt  = 0
    var i int = 0
    fmt.Println("result", i == j) //mismatched types int and MyInt
}
