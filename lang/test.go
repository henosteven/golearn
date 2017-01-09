package main

import (
    "fmt"
    "time"
)

type Vertex struct {
    x int
    y int
}

func sum(x int, y int) int {
    return x + y
}

func swap(a, b string) (string, string) {
    return b, a
}

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    
    v := Vertex{1, 2}
    fmt.Println(v)
    var p *Vertex
    p = &v
    p.x = 8
    fmt.Println(v)

    fmt.Println("hello world")
    fmt.Println("the time is:", time.Now())
    fmt.Println("the sum is :", sum(3, 4))

    a, b := swap("hello", "world")
    fmt.Println("swap: ", a, b)
    fmt.Println(split(17))

    c := []int{1, 2, 3}
    d := c
    d[0] = 9
    fmt.Println(c) -- 9 2 3
    fmt.Println(d) -- 9 2 3 
}
