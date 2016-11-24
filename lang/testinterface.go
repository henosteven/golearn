package main

import (
    "fmt"
)

type ishow interface {
    showname()
    showage()
    showany(str string)
}

type iany interface{}

type user struct {
    name string
    age int
}

func (u user) showname() {
    fmt.Println(u.name)
}

func (u user)showage() {
    fmt.Println(u.age)
} 

func (u user)showany(str string) {
    fmt.Println(str)
}

type student struct {
    name string
    age int
}

func (u student) showname() {
    fmt.Println(u.name)
}

func (u student)showage() {
    fmt.Println(u.age)
} 

func (u student)showany(str string) {
    fmt.Println(str)
}

func main() {
    var name iany
    name = "hello"
    fmt.Println(name, name.(string))

    f, ok := name.(string)
    fmt.Println(f, ok)

    var heno ishow = user{"jinjing", 27}
    heno.showname()
    heno.showage()
    heno.showany("hello world")
    fmt.Printf("%v %T\n", heno, heno)

    var pang ishow = student{"xiaopang", 27}
    pang.showname()
    pang.showage()
    pang.showany("hello world")
    fmt.Printf("%v %T\n", pang, pang)


    switch pang.(type) {
        case student:
            fmt.Println("type-student")
            break
        default:
            fmt.Println("not-match")
    }
    
    //cannot use "jinjing" (type string) as type ishow in assignment:
    //string does not implement ishow (missing showage method)
    //var hao ishow = "jinjing"
}
