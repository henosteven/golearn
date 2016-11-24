package main

import (
    "fmt"
)

func increScore() func() int {
    var a int = 0
    return func() int {
        a = a + 1
        return a
    }
}

func main() {
    f := increScore()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
