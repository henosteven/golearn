package main

import (
    "fmt"
)

func inputch(ch chan int) {
    ch <- 12
    ch <- 11
    close(ch)
}

func main() {
    ch := make(chan int)
    go inputch(ch)
    for v := range ch {
        fmt.Println(v)
    }
}
