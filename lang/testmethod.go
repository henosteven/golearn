package main

import (
    "fmt"
)

type user struct {
    name string
    age int
}

func (u *user)test() {
    u.name = "jinjing"
    u.age = 27
}

func change( u *user, name string) {
    u.name = name
}

func main() {
    var u user
    u.test()
    fmt.Println(u)
    change(&u, "heno")
    fmt.Println(u)
}
